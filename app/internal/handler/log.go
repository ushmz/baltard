package handler

import (
	"fmt"
	"net/http"

	"ratri/internal/domain/model"
	"ratri/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Log struct {
	usecase usecase.Log
}

func NewLogHandler(log usecase.Log) *Log {
	return &Log{usecase: log}
}

// CreateTaskTimeLog : Create task time log. Table name is `behacior_logs`.
// Id create_task_time_log
// Summary Store task time log
// Description Create task time log with value in the request. If key (user_id and task_id) is depulicated, update `time` value instead of creating new record.
// Accept json
// Produce json
// Param param body model.TaskTimeLogParamWithTime true "Log parameter"
// Success 200
// Failure 400
// Failure 500
// Router /v1/logs/time [POST]
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

// CumulateTaskTimeLog : Create task time log. Task time is counted by cumulating requests that should be sended once/sec.
// @Id cumulate_task_time_log
// @Summary Store task time log
// @Description Create task time log. Task time is measured by cumulating number of requests that should be sended once/sec.
// @Accept json
// @Produce json
// @Param param body model.TaskTimeLogParam true "Log parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/logs/time [POST]
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

// CreateSerpClickLog : Create click log.
// @Id create_serp_click_log
// @Summary Store SERP click log
// @Description Create click log in SERP.
// @Accept json
// @Produce json
// @Param param body model.SearchPageClickLogParam true "Log parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/logs/click [POST]
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

// StoreSearchSeeion : Store search session log.
// @Id store_search_session
// @Summary Store search session log
// @Description Store search session log that is consists of task start(User presses the "Start searching for the task" button) and end(User presses the "Submit answer" button) time.
// @Accept json
// @Produce json
// @Param param body model.SearchSession true "Log parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/logs/session [POST]
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
