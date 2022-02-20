package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"ratri/src/domain/model"
	"ratri/src/usecase"

	"github.com/labstack/echo/v4"
)

type Task struct {
	usecase usecase.TaskUsecase
}

func NewTaskHandler(task usecase.TaskUsecase) *Task {
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
			Message: "Parameter `id` must be number",
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
	p := model.Answer{}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: fmt.Sprintf("Invalid request body : %v", err),
		})
	}

	err := t.usecase.CreateTaskAnswer(&p)
	// Execute query.
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to submit answer.",
		})
	}

	return c.NoContent(http.StatusCreated)
}
