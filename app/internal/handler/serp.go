package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"ratri/internal/domain/model"
	"ratri/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Serp struct {
	usecase usecase.Serp
}

func NewSerpHandler(serp usecase.Serp) *Serp {
	return &Serp{usecase: serp}
}

// FetchSerpWithDistributionByID : Return search result pages with similarweb information (such as icon)
// @Id fetch_serp_with_distribution_by_id
// @Summary Returns json which have a list of search results with data for RatioUI.
// @Description Returns json which have a list of search results with tracked pages of Similarweb top 2000 and its category.
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param offset query int true "Number of offset to display"
// @Param top query int false "Number of categories to display"
// @Success 200 {object} []model.SerpWithRatio
// @Failure 400
// @Failure 500
// @Router /v1/serp/{taskId}/ratio [GET]
func (s *Serp) FetchSerpWithDistributionByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Parameter `taskId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")
	// offset : Parse string value to int value.
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		msg := model.ErrorMessage{
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
		msg := model.ErrorMessage{
			Message: "Parameter `top` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	serp, err := s.usecase.FetchSerpWithRatio(task, offset, top)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database execution error: Failed to fetch relations: " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpWithIconByID : Return search result pages with similarweb information (such as icon)
// @Id fetch_serp_with_icon_by_id
// @Summary Returns json which have a list of search results with data for IconUI.
// @Description Returns json which have a list of search results with tracked pages of Similarweb top 2000 pages and its favicon URL.
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param offset query int true "Number of offset to display"
// @Param top query int false "Number of icons to display"
// @Success 200 {object} []model.SerpWithIcon
// @Failure 400
// @Failure 500
// @Router /v1/serp/{taskId}/icon [GET]
func (s *Serp) FetchSerpWithIconByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		fmt.Println(err)
		msg := model.ErrorMessage{
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
		msg := model.ErrorMessage{
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
		msg := model.ErrorMessage{
			Message: "Parameter `top` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	serp, err := s.usecase.FetchSerpWithIcon(task, offset, top)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database execution error: Failed to fetch relations : " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpByID : Return search result pages by task id .
// @Id fetch_serp_by_id
// @Summary Returns json which have a list of search results.
// @Description Returns json which have a list of search results with no additional data.
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param offset query int true "Number of offset to display"
// @Success 200 {object} []model.SearchPage
// @Failure 400
// @Failure 500
// @Router /v1/serp/{taskId} [GET]
func (s *Serp) FetchSerpByID(c echo.Context) error {
	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Parameter `taskId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// offsetstr : Get offset from query parameter.
	offsetstr := c.QueryParam("offset")
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Parameter `offset` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	serp, err := s.usecase.FetchSerp(task, offset)

	return c.JSON(http.StatusOK, serp)
}
