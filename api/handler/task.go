package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ymmt3-lab/koolhaas/backend/api/models"
)

// FetchTaskInfo : Fetch task info by task id
func (h *Handler) FetchTaskInfo(c echo.Context) error {

	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `taskId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// Fetch task information by task Id
	ti, err := h.Task.FetchTaskInfo(task)
	if err != nil {
		if err == sql.ErrNoRows {
			// Unreachable code block
			c.Echo().Logger.Infof("TaskId %v not found", taskId)
			return c.NoContent(http.StatusNotFound)
		}
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, ti)
}

// SubmitTaskAnswer : Submit task answer
func (h *Handler) SubmitTaskAnswer(c echo.Context) error {
	// answer : Bind request body to struct
	answer := new(models.TaskAnswer)
	if err := c.Bind(answer); err != nil {
		c.Echo().Logger.Errorf("Error. Invalid request body : %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	err := h.Answer.SubmitTaskAnswer(answer)
	// Execute query.
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Failed to submit answer.",
		})
	}

	return c.NoContent(http.StatusCreated)
}
