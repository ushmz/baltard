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
			serp_sim2000_relation_top10
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
		serp_sim2000_relation_top10
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
			id,
			title,
			url,
			icon_path
		FROM
			similarweb_2000_pages
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
		for _, v := range similarIds {
			leaks = append(leaks, similarsMap[v])
		}
		response = append(response, models.Serp{
			PageId:  page.PageId,
			Title:   page.Title,
			Url:     page.Url,
			Snippet: page.Snippet,
			Leaks:   leaks,
		})
	}
	c.Echo().Logger.Info(response)
	return c.JSON(http.StatusOK, response)
}
