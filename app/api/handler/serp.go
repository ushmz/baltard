package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"baltard/api/model"
	"baltard/api/service"

	"github.com/labstack/echo"
)

type Serp struct {
	service service.Serp
}

func NewSerpHandler(serpService service.Serp) *Serp {
	return &Serp{service: serpService}
}

// FetchSerpWithDistributionByID : Return search result pages with similarweb information (such as icon)
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

	serp, err := s.service.FetchSerpWithRatio(task, offset, top)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database execution error: Failed to fetch relations: " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpWithIconByID : Return search result pages with similarweb information (such as icon)
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

	serp, err := s.service.FetchSerpWithIcon(task, offset, top)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database execution error: Failed to fetch relations : " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpByID : Return search result pages by task id .
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

	serp, err := s.service.FetchSerp(task, offset)

	return c.JSON(http.StatusOK, serp)
}
