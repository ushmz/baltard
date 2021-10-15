package handler

import (
	"fmt"
	"net/http"

	"ratri/internal/domain/model"
	"ratri/internal/usecase"

	"github.com/labstack/echo"
)

type Log struct {
	usecase usecase.Log
}

func NewLogHandler(log usecase.Log) *Log {
	return &Log{usecase: log}
}

// CreateTaskTimeLog : Create task time log. Table name is `behacior_logs`.
// Create log one record by user id, if its id is depulicated, update record instead create new record.
func (l *Log) CreateTaskTimeLog(c echo.Context) error {
	// param : Bind request body to struct.
	param := new(model.TaskTimeLogParamWithTime)
	if err := c.Bind(param); err != nil {
		c.Echo().Logger.Errorf("Cannot bind request body to struct : %v", err)
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.StoreTaskTimeLog(param)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

func (l *Log) CumulateTaskTimeLog(c echo.Context) error {
	// param : Bind request body to struct.
	param := new(model.TaskTimeLogParam)
	if err := c.Bind(param); err != nil {
		c.Echo().Logger.Errorf("Cannot bind request body to struct : %v", err)
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.CumulateTaskTimeLog(param)
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
	// param : Bind request body to struct.
	param := new(model.SearchPageClickLogParam)
	if err := c.Bind(param); err != nil {
		c.Echo().Logger.Errorf("Failed to bind request body : %v", err)
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.StoreSerpClickLog(param)
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
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.StoreSearchSeeion(s)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}
