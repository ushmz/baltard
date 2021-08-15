package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"baltard/api/models"

	"github.com/labstack/echo"
)

// FetchSerp : Wrapper function
// func (h *Handler) FetchSerp(c echo.Context) error {}

func (h *Handler) FetchSerpWithDistributionByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `taskId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")
	// offset : Parse string value to int value.
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `offset` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// topstr : Return this number of top category.
	topstr := c.QueryParam("top")
	if topstr == "" {
		topstr = "3"
	}
	// top : Parse string value to int value.
	top, err := strconv.Atoi(topstr)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `top` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	swd, err := h.Serp.FetchSerpWithDistributionByID(task, offset, top)

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
				Distribution: []models.CategoryDistribution{},
			}
		}

		if v.Category != "" {
			if v.CategoryRank <= top {
				tempSerp := serpMap[v.PageId]
				tempSerp.Distribution = append(tempSerp.Distribution, models.CategoryDistribution{
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

// FetchSerpWithIconByID : Return search result pages with similarweb information (such as icon)
func (h *Handler) FetchSerpWithIconByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Parameter `taskId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")
	// offset : Parse offset string to int value.
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		fmt.Println(err)
		msg := models.ErrorMessage{
			Message: "Parameter `offset` must be number",
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
			Message: "Parameter `top` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	swi, err := h.Serp.FetchSerpWithIconByID(task, offset, top)
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
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(v.Leaks), func(i, j int) { v.Leaks[i], v.Leaks[j] = v.Leaks[j], v.Leaks[i] })

		serp = append(serp, v)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpByID : Return search result pages by task id .
func (h *Handler) FetchSerpByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `taskId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `offset` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	s, err := h.Serp.FetchSerpByID(task, offset)

	return c.JSON(http.StatusOK, s)
}
