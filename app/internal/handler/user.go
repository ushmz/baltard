package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"ratri/internal/domain/model"
	"ratri/internal/usecase"

	"github.com/labstack/echo/v4"
)

type User struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(user usecase.UserUsecase) *User {
	return &User{usecase: user}
}

// CreateUser : Register new user with crowd-sourcing service ID
// @Id create_user
// @Summary Register new user.
// @Description Register user using crowd-sourcing service ID and allocate task at the same time. If external service ID is already exists, re-allocate task and return.
// @Accept json
// @Produce json
// @Param param body model.UserParam true "User parameter"
// @Success 200 {object} model.User
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /users [POST]
func (u *User) CreateUser(c echo.Context) error {
	// u : Request body struct
	p := model.UserParam{}
	// Bind request body parameters to struct
	if err := c.Bind(&p); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// Verbose
	var user model.User

	// exist : Given uid is already exist or not
	eu, exist, err := u.usecase.FindByUid(p.Uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to detect user existence.",
		})
	}

	if exist {
		user = eu
	} else {
		nu, err := u.usecase.CreateUser(p.Uid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
				Message: "Failed to create new user.",
			})
		}
		user = nu
	}

	info, err := u.usecase.AllocateTask()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to allocate task.",
		})
	}

	return c.JSON(http.StatusOK, model.UserResponse{
		Exist:       exist,
		UserId:      user.Id,
		Secret:      user.Secret,
		TaskIds:     info.TaskIds,
		ConditionId: info.ConditionId,
		GroupId:     info.GroupId,
	})
}

// GetCompletionCode : Get users task completion code.
// @Id get_completion_code
// @Summary Get completion code.
// @Description Get unique task completion code.
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {int} int "Completion code"
// @Failure 400 "Error with message"
// @Failure 500 "Error with message"
// @Router /v1/users/code/{id} [GET]
func (u *User) GetCompletionCode(c echo.Context) error {
	// id : Get id from path parameter
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		msg := model.ErrorMessage{
			Message: "Parameter `userId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// Fetch completion code by uid from DB
	code, err := u.usecase.GetCompletionCode(userId)
	if err != nil {
		// If given uid not found in DB
		if err == sql.ErrNoRows {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, code)
}
