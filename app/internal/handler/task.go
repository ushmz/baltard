package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"ratri/internal/domain/model"
	"ratri/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Task struct {
	usecase usecase.Task
}

func NewTaskHandler(task usecase.Task) *Task {
	return &Task{usecase: task}
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

	// taskId : Get task Id from path parameter.
	taskId := c.Param("id")
	task, err := strconv.Atoi(taskId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Parameter `taskId` must be number",
		})
	}

	// Fetch task information by task Id
	ti, err := t.usecase.FetchTaskInfo(task)
	if err != nil {
		if errors.Is(err, model.NoSuchDataError{}) {
			return c.JSON(http.StatusNotFound, model.ErrorMessage{
				Message: fmt.Sprintf("TaskId %v not found", taskId),
			})
		}
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
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
	// answer : Bind request body to struct
	answer := new(model.Answer)
	if err := c.Bind(answer); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: fmt.Sprintf("Invalid request body : %v", err),
		})
	}

	err := t.usecase.CreateTaskAnswer(answer)
	// Execute query.
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to submit answer.",
		})
	}

	return c.NoContent(http.StatusCreated)
}
