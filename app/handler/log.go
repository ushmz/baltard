package handler

import (
	"fmt"
	"net/http"

	"ratri/domain/model"
	"ratri/domain/store"
	"ratri/usecase"

	"github.com/labstack/echo/v4"
)

// Log : Implemention of log handler
type Log struct {
	usecase usecase.LogUsecase
}

// NewLogHandler : Return new log handler
func NewLogHandler(log usecase.LogUsecase) *Log {
	return &Log{usecase: log}
}

// CumulateSerpDwellTime : Create task time log. Task time is counted by cumulating requests that should be sended once/sec.
// @Id cumulate_task_time_log
// @Summary Store task time log
// @Description Create task time log. Task time is measured by cumulating number of requests that should be sended once/sec.
// @Accept json
// @Produce json
// @Param param body model.SerpDwellLogParam true "Log parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/logs/serp [POST]
func (l *Log) CumulateSerpDwellTime(c echo.Context) error {
	p := model.SerpDwellTimeLogParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.CumulateSerpDwellTime(p)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

// FileExportParam : Request parameters for exporting endpoints
type FileExportParam struct {
	Header   bool           `json:"header" query:"header"`
	FileType store.FileType `json:"type" query:"type"`
}

// ExportSerpDwellTime : Export all task time log.
// @Id export_task_time_log
// @Summary Export task time log
// @Description Export all task time log.
// @Accept json
// @Produce text/csv text/tab-separated-values
// @Param param query FileExportParam true "Export parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/logs/serp/export [GET]
func (l *Log) ExportSerpDwellTime(c echo.Context) error {
	p := FileExportParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	b, err := l.usecase.ExportPageDwellTimeLog(p.Header, p.FileType)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Failed to export data.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	if p.FileType == store.TSV {
		c.Response().Header().Set("Content-Type", "text/tab-separated-values")
	} else {
		c.Response().Header().Set("Content-Type", "text/csv")
	}

	return c.JSONBlob(http.StatusOK, b.Bytes())
}

// CumulatePageDwellTime : Create page viewing time log. Dwell time is counted by cumulating requests that should be sended once/sec.
// @Id cumulate_page_viewing_time
// @Summary Store page viewing time log
// @Description Create page viewing time log. Dwell time is measured by cumulating number of requests that should be sended once/sec.
// @Accept json
// @Produce json
// @Param param body model.PageDwellLogParam true "Log parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/logs/pageview [POST]
func (l *Log) CumulatePageDwellTime(c echo.Context) error {
	p := model.PageDwellTimeLogParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.CumulatePageDwellTime(p)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

// ExportPageDwellTime : Export all page dwell time logs to file
func (l *Log) ExportPageDwellTime(c echo.Context) error {
	p := FileExportParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	b, err := l.usecase.ExportPageDwellTimeLog(p.Header, p.FileType)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Failed to export log.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	if p.FileType == store.TSV {
		c.Response().Header().Set("Content-Type", "text/tab-separated-values")
	} else {
		c.Response().Header().Set("Content-Type", "text/csv")
	}

	return c.JSONBlob(http.StatusOK, b.Bytes())
}

// CreateSerpEventLog : Create click log.
// @Id create_serp_click_log
// @Summary Store SERP click log
// @Description Create click log in SERP.
// @Accept json
// @Produce json
// @Param param body model.SearchPageEventLogParam true "Log parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/logs/click [POST]
func (l *Log) CreateSerpEventLog(c echo.Context) error {
	p := model.SearchPageEventLogParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.StoreSerpEventLog(p)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

// ExportSerpEventLog : Export all SERP event logs to file
func (l *Log) ExportSerpEventLog(c echo.Context) error {
	p := FileExportParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: fmt.Sprintf("Cannot bind request body : %v", err),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	b, err := l.usecase.ExportSerpEventLog(p.Header, p.FileType)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Failed to export log.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	if p.FileType == store.TSV {
		c.Response().Header().Set("Content-Type", "text/tab-separated-value")
	} else {
		c.Response().Header().Set("Content-Type", "csv")
	}

	return c.JSONBlob(http.StatusOK, b.Bytes())
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
	p := model.SearchSessionParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: "Invalid request body.",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	err := l.usecase.StoreSearchSeeion(p)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Database Execution error.",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.NoContent(http.StatusCreated)
}

// ExportSearchSeeion : Export all search sessions to file
func (l *Log) ExportSearchSeeion(c echo.Context) error {
	p := FileExportParam{}
	if err := c.Bind(&p); err != nil {
		msg := model.ErrorMessage{
			Message: "Invalid request body.",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	b, err := l.usecase.ExportSearchSeeion(p.Header, p.FileType)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Failed to export log",
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	if p.FileType == store.TSV {
		c.Response().Header().Set("Content-Type", "text/tab-separated-value")
	} else {
		c.Response().Header().Set("Content-Type", "text/tab-separated-value")
	}

	return c.JSONBlob(http.StatusOK, b.Bytes())
}
