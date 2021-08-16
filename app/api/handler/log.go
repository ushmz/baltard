package handler

import (
	"net/http"

	"baltard/api/model"
	"baltard/api/service"

	"github.com/labstack/echo"
)

type Log struct {
	service service.Log
}

func NewLogHandler(logService service.Log) *Log {
	return &Log{service: logService}
}

// CreateTaskTimeLog : Create task time log. Table name is `behacior_logs`.
// Create log one record by user id, if its id is depulicated, update record instead create new record.
func (l *Log) CreateTaskTimeLog(c echo.Context) error {
	// l : Bind request body to struct.
	param := new(model.TaskTimeLogParam)
	if err := c.Bind(param); err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	err := l.service.CreateTaskTimeLog(param)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

func (l *Log) CreateSerpClickLog(c echo.Context) error {
	// p : Bind request body to struct.
	param := new(model.SearchPageClickLogParam)
	if err := c.Bind(param); err != nil {
		c.Echo().Logger.Errorf("Failed to bind request body : %v", err)
		msg := model.ErrorMessage{
			Message: "Failed to bind request body.",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.service.CreateSerpClickLog(param)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

func (l *Log) StoreSearchSeeion(c echo.Context) error {
	s := new(model.SearchSession)
	if err := c.Bind(s); err != nil {
		c.Echo().Logger.Errorf("Invalid request body : %v", err)
		msg := model.ErrorMessage{
			Message: "Invalid request body.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	err := l.service.StoreSearchSeeion(s)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}
