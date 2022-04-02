package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"ratri/domain/model"
	"ratri/usecase"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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

	// exist : Given uid is already exist or not
	user, err := u.usecase.FindByUid(p.Uid)
	if err != nil {
		switch errors.Cause(err) {
		case model.ErrNoSuchData:
			user, err = u.usecase.CreateUser(p.Uid)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
					Message: "Failed to create new user.",
				})
			}
		default:
			return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
				Message: "Failed to detect user existence.",
			})
		}
	}

	info, err := u.usecase.AllocateTask()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to allocate task.",
		})
	}

	cc := createCookie("exp-condition", fmt.Sprint(info.ConditionId))
	c.SetCookie(cc)
	gc := createCookie("exp-group", fmt.Sprint(info.GroupId))
	c.SetCookie(gc)
	ftc := createCookie("exp-1-task", fmt.Sprint(info.TaskIds[0]))
	c.SetCookie(ftc)
	stc := createCookie("exp-2-task", fmt.Sprint(info.TaskIds[1]))
	c.SetCookie(stc)

	st := createCookie("exp-token", user.Token)
	c.SetCookie(st)

	return c.JSON(http.StatusOK, model.UserResponse{
		Token:       user.Token,
		UserId:      user.Id,
		TaskIds:     info.TaskIds,
		ConditionId: info.ConditionId,
		GroupId:     info.GroupId,
	})
}

func createCookie(name string, val string) *http.Cookie {
	c := new(http.Cookie)
	c.HttpOnly = true
	c.Secure = true
	c.Path = "/"
	c.Expires = time.Now().Add(1 * time.Hour)
	c.Name = name
	c.Value = val
	return c
}

type CreateSessionParameter struct {
	IDToken string `json:"token"`
}

func (u *User) CreateSession(c echo.Context) error {
	if u == nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	p := new(CreateSessionParameter)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	cval, err := u.usecase.CreateSession(p.IDToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	sc := createCookie("session", cval)
	c.SetCookie(sc)
	return c.NoContent(http.StatusOK)
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
