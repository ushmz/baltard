package handler

import (
	"database/sql"
	"net/http"

	"koolhaas/backend/models"

	"github.com/labstack/echo"
)

// FetchTaskInfo : Fetch task info by task id
func (h *Handler) FetchTaskInfo(c echo.Context) error {

	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")

	// task : Task information.
	task := []models.Task{}
	// Fetch task information by task Id
	err := h.DB.Select(&task, `
		SELECT
			id,
			query,
			title,
			description,
			search_url
		FROM
			tasks
		WHERE
			id = ?
		`, taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Unreachable code block
			c.Echo().Logger.Infof("TaskId %v not found", taskId)
			return c.NoContent(http.StatusNotFound)
		}
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, task)
}

// SubmitTaskAnswer : Submit task answer
func (h *Handler) SubmitTaskAnswer(c echo.Context) error {
	// answer : Bind request body to struct
	answer := new(models.TaskAnswer)
	if err := c.Bind(answer); err != nil {
		c.Echo().Logger.Errorf("Error. Invalid request body : %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	// query : Task answer insert query with named format characters.
	query := `
		INSERT INTO
			answers (
				uid,
				task_id,
				condition_id,
				author_id,
				answer,
				reason
			)
		VALUES (
			:uid,
			:task_id,
			:condition_id,
			2,
			:answer,
			:reason
		)`

	// vals : Parameters of SQL query.
	vals := &answer

	// Execute query.
	if _, err := h.DB.NamedExec(query, vals); err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Failed to submit answer.",
		})
	}

	return c.NoContent(http.StatusCreated)
}
