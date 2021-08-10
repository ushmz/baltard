package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymmt3-lab/koolhaas/backend/api/models"
)

// CreateTaskTimeLog : Create task time log. Table name is `behacior_logs`.
// Create log one record by user id, if its id is depulicated, update record instead create new record.
func (h *Handler) CreateTaskTimeLog(c echo.Context) error {
	// l : Bind request body to struct.
	l := new(models.TaskTimeLogParam)
	if err := c.Bind(l); err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := models.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	err := h.Log.CreateTaskTimeLog(l)
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
	if err := c.Bind(p); err != nil {
		c.Echo().Logger.Errorf("Failed to bind request body : %v", err)
		msg := models.ErrorMessage{
			Message: "Failed to bind request body.",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := h.Log.CreateSerpClickLog(p)
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

	err := h.Log.StoreSearchSeeion(s)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := models.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}
