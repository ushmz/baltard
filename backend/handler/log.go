package handler

import (
	"net/http"

	"koolhaas/backend/models"

	"github.com/labstack/echo"
)

// CreateTaskTimeLog : Create task time log. Table name is `behacior_logs`.
// Create log one record by user id, if its id is depulicated, update record instead create new record.
func (h *Handler) CreateTaskTimeLog(c echo.Context) error {
	// l : Bind request body to struct.
	l := new(models.TaskTimeLogParam)
	var err error
	if err = c.Bind(l); err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := models.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// query : Task log insert query with format characters.
	query := `
		INSERT INTO
			behavior_logs (
				user_id,
				task_id,
				time_on_page,
				condition_id
			)
		VALUES (
			:user_id, 
			:task_id, 
			:time_on_page, 
			:condition_id
		)
		ON DUPLICATE
			KEY UPDATE
				time_on_page = :time_on_page, 
				updated_at = CURRENT_TIMESTAMP`

	// Execute query
	_, err = h.DB.NamedExec(query, l)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := models.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) CreateSerpClickLog(c echo.Context) error {
	// p : Bind request body to struct.
	p := new(models.SearchPageClickLogParam)
	var err error
	if err = c.Bind(p); err != nil {
		c.Echo().Logger.Errorf("Failed to bind request body : %v", err)
		msg := models.ErrorMessage{
			Message: "Failed to bind request body.",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// query : Task log insert query.
	query := `
		INSERT INTO
			behavior_logs_click (
				user_id,
				task_id,
				condition_id,
				time_on_page,
				serp_page,
				serp_rank,
				is_visible
			)
		VALUES (
			:user_id, 
			:task_id, 
			:condition_id,
			:time_on_page, 
			:serp_page,
			:serp_rank,
			:is_visible
		)`

	// Execute SQL query.
	_, err = h.DB.NamedExec(query, p)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := models.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) StoreSearchSeeion(c echo.Context) error {
	s := new(models.SearchSession)
	if err := c.Bind(s); err != nil {
		c.Echo().Logger.Errorf("Invalid request body : %v", err)
		msg := models.ErrorMessage{
			Message: "Invalid request body.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	_, err := h.DB.NamedExec(`
		INSERT INTO
			search_session(
				user_id,
				task_id,
				condition_id
			)
		VALUES (
			:user_id,
			:task_id,
			:condition_id
		)
		ON DUPLICATE KEY
			UPDATE
				ended_at = CURRENT_TIMESTAMP
		`, s)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := models.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}
