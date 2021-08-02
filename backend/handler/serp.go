package handler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/ymmt3-lab/koolhaas/backend/models"
)

// FetchSerp : Wrapper function
// func (h *Handler) FetchSerp(c echo.Context) error {}

func (h *Handler) FetchSerpWithDistributionByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")

	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")
	// offset : Parse string value to int value.
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `offset` must be int value",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// topstr : Return this number of top category.
	topstr := c.QueryParam("top")
	// If value is not specified, set default value `3`
	if topstr == "" {
		topstr = "3"
	}
	// top : Parse string value to int value.
	top, err := strconv.Atoi(topstr)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `top` must be int value",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	swd := []models.SerpWithDistributionQueryResult{}
	err = h.DB.Select(&swd, `
		SELECT
			relation_count.id,
			relation_count.title,
			relation_count.url,
			relation_count.snippet,
			relation_count.category,
			ROW_NUMBER() OVER(
				PARTITION BY relation_count.id
				ORDER BY relation_count.category_count
				DESC
			) category_rank,
			relation_count.category_count,
			relation_count.similarweb_count,
			relation_count.category_count / relation_count.similarweb_count category_distribution
		FROM (
			SELECT DISTINCT
				search_pages.id,
				search_pages.title,
				search_pages.url,
				search_pages.snippet,
				similarweb_categories.category,
				COUNT(*) OVER(PARTITION BY search_pages.id) similarweb_count,
				COUNT(*) OVER(PARTITION BY search_pages.id, similarweb_categories.category) category_count
			FROM ( SELECT
				page_id,
				similarweb_id,
				idf
			FROM
				search_page_similarweb_relation
			WHERE
				page_id IN (SELECT * FROM (
					SELECT
						page_id
					FROM
						search_page_similarweb_relation
					WHERE
						task_id = ?
					GROUP BY
						page_id
					LIMIT ?, 10
					) as result)
				) as relation
				JOIN
					search_pages
				ON
					relation.page_id = search_pages.id
				JOIN
					similarweb_pages
				ON
					relation.similarweb_id = similarweb_pages.id
				JOIN
					similarweb_categories
				ON
					similarweb_pages.category = similarweb_categories.id
			) as relation_count
	`, taskId, offset*10)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to fetch relations: " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]models.SerpWithDistribution{}

	// serp : Return struct of this method
	serp := []models.SerpWithDistribution{}

	for _, v := range swd {
		if _, ok := serpMap[v.PageId]; !ok {
			serpMap[v.PageId] = models.SerpWithDistribution{
				PageId:       v.PageId,
				Title:        v.Title,
				Url:          v.Url,
				Snippet:      v.Snippet,
				Total:        v.SimilarwebCount,
				Distribution: []models.SimilarwebDistribution{},
			}
		}

		if v.Category != "" {
			if v.CategoryRank <= top {
				tempSerp := serpMap[v.PageId]
				tempSerp.Distribution = append(tempSerp.Distribution, models.SimilarwebDistribution{
					Category:   v.Category,
					Count:      v.CategoryCount,
					Percentage: v.CategoryDistribution,
				})
				serpMap[v.PageId] = tempSerp
			}
		}
	}

	for _, v := range serpMap {
		serp = append(serp, v)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpWithIconByID : Return search result pages with similarweb icon
func (h *Handler) FetchSerpWithIconByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")

	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")
	// offset : Parse offset string to int value.
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Parameter `offset` must be int value",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// topstr : Return this number of top category.
	topstr := c.QueryParam("top")
	// If value is not specified, set default value `3`
	if topstr == "" {
		topstr = "10"
	}
	// top : Parse string value to int value.
	top, err := strconv.Atoi(topstr)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `top` must be int value",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// swi : SQL query result, serp with top 10 of largest similarweb idf
	swi := []models.SerpWithIconQueryResult{}
	err = h.DB.Select(&swi, `
		SELECT
			search_pages.id,
			search_pages.title,
			search_pages.url,
			search_pages.snippet,
			similarweb_pages.id similarweb_id,
			similarweb_pages.title similarweb_title,
			similarweb_pages.url similarweb_url,
			similarweb_pages.icon_path similarweb_icon,
			similarweb_categories.category similarweb_category
		FROM (
			SELECT
				page_id,
				similarweb_id,
				idf,
				ROW_NUMBER() OVER(PARTITION BY page_id ORDER BY idf DESC) idf_rank
			FROM
				search_page_similarweb_relation
			WHERE
				page_id IN (SELECT * FROM (
					SELECT
						page_id
					FROM
						search_page_similarweb_relation
					WHERE
						task_id = ?
					GROUP BY
						page_id
					LIMIT ?, 10
				) as result)
			) as relation
		JOIN
			search_pages
		ON
			relation.page_id = search_pages.id
		JOIN
			similarweb_pages
		ON
			relation.similarweb_id = similarweb_pages.id
		JOIN
			similarweb_categories
		ON
			similarweb_pages.category = similarweb_categories.id
		WHERE
			relation.idf_rank <= ?
		`, taskId, offset*10, top)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to fetch relations : " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]models.SerpWithIcon{}

	// serp : Return struct of this method
	serp := []models.SerpWithIcon{}

	for _, v := range swi {
		if _, ok := serpMap[v.PageId]; !ok {
			serpMap[v.PageId] = models.SerpWithIcon{
				PageId:  v.PageId,
				Title:   v.Title,
				Url:     v.Url,
				Snippet: v.Snippet,
				Leaks:   []models.SimilarwebPage{},
			}
		}

		if v.SimilarwebId != 0 {
			tempSerp := serpMap[v.PageId]
			tempSerp.Leaks = append(tempSerp.Leaks, models.SimilarwebPage{
				Id:       v.SimilarwebId,
				Title:    v.SimilarwebTitle,
				Url:      v.SimilarwebUrl,
				Icon:     v.SimilarwebIcon,
				Category: v.SimilarwebCategory,
			})
			serpMap[v.PageId] = tempSerp
		}
	}

	for _, v := range serpMap {
		serp = append(serp, v)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpByID : Return search result pages with probably leaked pages by task id .
func (h *Handler) FetchSerpByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")

	// offset : Parse offset string to int value.
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Parameter `offset` must be int value",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// rel : Relation with search page Ids and linked similarweb page Ids
	rel := []models.SerpSimilarwebRelation{}

	err = h.DB.Select(&rel, `
		SELECT
			page_id,
			similarweb_id
		FROM
			search_page_similarweb_relation
		WHERE
			task_id = ?
	`, taskId)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to fetch relations",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// relMap : Key: search page id, value: similarweb Ids
	relMap := map[int][]int{}
	// Convert relations array to map
	for _, v := range rel {
		relMap[v.PageId] = append(relMap[v.PageId], v.SimilarwebId)
	}

	// rng : Get page Id range with task Id
	var rng []models.SerpRange
	// Fetch search page Id range from DB.
	err = h.DB.Select(&rng, `
		SELECT
			min_id,
			max_id
		FROM
			page_id_range
		WHERE
			task_id = ?
	`, taskId)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to fetch page id range",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// maxOffset : Get max page number by stored pages.
	maxOffset := (rng[0].Max-rng[0].Min)/10.0 - 1
	// offset : Get min value with comparing maxOffset and requested offset value.
	offset = int(math.Min(float64(maxOffset), float64(offset)))
	// minId : Get minimum ID from offset value.
	// minId := rng[0].Min + 10*offset
	// maxId : Get maximum ID from offset value.
	// maxId := minId + 9

	var intrng []int
	err = h.DB.Select(&intrng, `
	SELECT
		page_id
	FROM
		search_page_similarweb_relation
	WHERE
		task_id = ?
	GROUP BY
		page_id
	LIMIT
		?, 10
	`, taskId, 10*offset)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to fetch search pages range",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// Generate SQL `in` query
	q, p, err := sqlx.In(`
		SELECT
			id,
			title,
			url,
			snippet
		FROM
			search_pages
		WHERE
			task_id = ?
		AND
			id IN (?)
	`, taskId, intrng)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to generate in query",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// serp : Search result pages.
	serp := []models.SearchPage{}
	// Fetch search pages from DB.
	if err = h.DB.Select(&serp, q, p...); err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to fetch search pages",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// serpMap : Map object to avoid deprication and easy to link with similarweb pages.
	serpMap := map[int]models.SearchPage{}
	// Convert search page array to map
	for _, v := range serp {
		serpMap[v.PageId] = v
	}

	// similars : Similarweb Ppge
	similars := []models.SimilarwebPage{}
	// Fetch all similarweb pages from DB.
	err = h.DB.Select(&similars, `
		SELECT
			similarweb_pages.id,
			similarweb_pages.title,
			similarweb_pages.url,
			similarweb_pages.icon_path,
			similarweb_categories.category
		FROM
			similarweb_pages
		JOIN
			similarweb_categories
		ON
			similarweb_pages.category = similarweb_categories.id
	`)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Database execution error: Failed to fetch similarweb pages",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// similarsMap : Map object to avoid deprication and easy to link with serp pages.
	similarsMap := map[int]models.SimilarwebPage{}
	// Convert similarweb array to map
	for _, v := range similars {
		similarsMap[v.Id] = v
	}

	// response : Search result page with similarweb pages.
	response := []models.Serp{}

	// Loop each search result page and link it with similarweb page.
	for pageId, page := range serpMap {
		// similarIds : similarweb Ids that linked with the search result page.
		similarIds := relMap[pageId]
		// leaks : similarweb pages that linked with the search result page.
		leaks := []models.SimilarwebPage{}

		percent := map[string]int{}
		categ := map[string]float64{}
		for _, v := range similarIds {
			leaks = append(leaks, similarsMap[v])
		}

		for _, v := range leaks {
			percent[v.Category] += 1
		}

		sum := 0
		for _, v := range percent {
			sum += v
		}

		for k, v := range percent {
			fmt.Println(k, v, sum)
			fmt.Println(sum, " = ", len(leaks))
			categ[k] = float64(v) / float64(sum)
		}

		response = append(response, models.Serp{
			PageId:       page.PageId,
			Title:        page.Title,
			Url:          page.Url,
			Snippet:      page.Snippet,
			Leaks:        leaks,
			Distribution: categ,
		})
	}
	c.Echo().Logger.Info(response)
	return c.JSON(http.StatusOK, response)
}
