package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymmt3-lab/koolhaas/backend/models"
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
			c.Echo().Logger.Infof("TaskId %v not found", taskId)
			return c.NoContent(http.StatusNotFound)
		}
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	c.Echo().Logger.Info(task)
	return c.JSON(http.StatusOK, task)
}

// SubmitTaskAnswer : Submit task answer
func (h *Handler) SubmitTaskAnswer(c echo.Context) error {
	// answer : Bind request body to struct
	answer := new(models.TaskAnswer)
	var err error
	if err = c.Bind(answer); err != nil {
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
			:author_id,
			:answer,
			:reason
		)`

	// vals : Parameters of SQL query.
	vals := &answer

	// Execute query.
	_, err = h.DB.NamedExec(query, vals)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, answer)
}
