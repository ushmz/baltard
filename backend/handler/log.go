package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymmt3-lab/koolhaas/backend/models"
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
		return c.JSON(http.StatusInternalServerError, msg)
	}

	// query : Task log insert query with format characters.
	query := `
		INSERT INTO
			behavior_logs (
				id,
				uid,
				time_on_page,
				url,
				task_id,
				condition_id
			)
		VALUES (
			:id, 
			:uid, 
			:time_on_page, 
			:url, 
			:task_id, 
			:condition_id
		)
		ON DUPLICATE
			KEY UPDATE
				uid = :uid, 
				time_on_page = :time_on_page, 
				url = :url, 
				task_id = :task_id, 
				condition_id = :condition_id, 
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
