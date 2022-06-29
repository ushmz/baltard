package handler

import (
	"errors"
	"fmt"
	"net/http"

	"ratri/domain/model"
	"ratri/usecase"

	"github.com/labstack/echo/v4"
)

// Task : Implemention of task handler
type Task struct {
	usecase usecase.TaskUsecase
}

// NewTaskHandler : Return new task handler
func NewTaskHandler(task usecase.TaskUsecase) *Task {
	return &Task{usecase: task}
}

// FetchTaskInfoParams : Request parameters for fetch task info
type FetchTaskInfoParams struct {
	ID int `json:"id" param:"id"`
}

// FetchTaskInfo : Fetch task info by task id
// @Id fetch_task_info
// @Summary Fetch task information.
// @Description Fetch task information by requeted task ID.
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} model.Task
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/task/{id} [GET]
func (t *Task) FetchTaskInfo(c echo.Context) error {
	if t == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, model.ErrNilReceiver)
	}

	p := FetchTaskInfoParams{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrWithMessage{
				error: fmt.Errorf("Cannot bind request body: %w", err),
				Why:   "Parameter `id` is needed",
			},
		)
	}

	// Fetch task information by task Id
	ti, err := t.usecase.FetchTaskInfo(p.ID)
	if err != nil {
		if errors.Is(err, model.ErrNoSuchData) {
			msg := fmt.Sprintf("Task with ID(%d) is not found", p.ID)
			return echo.NewHTTPError(
				http.StatusNotFound,
				ErrWithMessage{
					error: fmt.Errorf("%s: %w", msg, err),
					Why:   msg,
				},
			)
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{
				error: fmt.Errorf("Try to get search result: %w", err),
				Why:   "Something went wrong",
			},
		)
	}

	return c.JSON(http.StatusOK, ti)
}

// SubmitTaskAnswer : Submit task answer
// @Id submit_task_answer
// @Summary Submit task answer.
// @Description Submit task answer.
// @Accept json
// @Produce json
// @Param param body model.Answer true "Answer parameter"
// @Success 200
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/task/answer [POST]
func (t *Task) SubmitTaskAnswer(c echo.Context) error {
	if t == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, model.ErrNilReceiver)
	}

	p := model.Answer{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrWithMessage{
				error: fmt.Errorf("Invalid request body: %w", err),
				Why:   "Invalid request body",
			},
		)
	}

	if err := t.usecase.CreateTaskAnswer(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{
				error: fmt.Errorf("Try to create answer: %w", err),
				Why:   "Something went wrong",
			},
		)
	}

	return c.NoContent(http.StatusCreated)
}
