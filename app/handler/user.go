package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"ratri/config"
	"ratri/domain/model"
	"ratri/usecase"

	"github.com/labstack/echo/v4"
)

// User : Implemention of user handler
type User struct {
	usecase usecase.UserUsecase
}

// NewUserHandler : Return new user handler
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
	if u == nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{
				error: fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver),
				Why:   "",
			},
		)
	}

	// u : Request body struct
	p := model.UserParam{}
	// Bind request body parameters to struct
	if err := c.Bind(&p); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// exist : Given uid is already exist or not
	user, err := u.usecase.FindByUID(p.UID)
	if err != nil {
		if errors.Is(err, model.ErrNoSuchData) {
			user, err = u.usecase.CreateUser(p.UID)
			if err != nil {
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					ErrWithMessage{
						error: err,
						Why:   "Try to create new user.",
					},
				)
			}
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{
				error: err,
				Why:   "Try to detect user existence.",
			},
		)
	}

	info, err := u.usecase.AllocateTask()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, ErrWithMessage{
				error: err,
				Why:   "Try to allocate task.",
			},
		)
	}

	return c.JSON(http.StatusOK, model.UserResponse{
		Token:       user.Token,
		UserID:      user.ID,
		TaskIDs:     info.TaskIDs,
		ConditionID: info.ConditionID,
		GroupID:     info.GroupID,
	})
}

func createCookie(name string, val string) *http.Cookie {
	conf := config.GetConfig()
	isProd := conf.GetString("env") != "dev"
	c := new(http.Cookie)
	c.HttpOnly = true
	c.Secure = isProd
	c.Path = "/"
	c.Expires = time.Now().Add(1 * time.Hour)
	c.Name = name
	c.Value = val
	return c
}

// CreateSessionParameter : Request parameters for `CreateSession`
type CreateSessionParameter struct {
	IDToken string `json:"token"`
}

// CreateSession : Generate session token by idToken
func (u *User) CreateSession(c echo.Context) error {
	if u == nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{error: model.ErrNilReceiver, Why: "Something went wrong with Server"},
		)
	}

	p := new(CreateSessionParameter)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrWithMessage{
				error: err,
				Why:   "Invalid request body",
			},
		)
	}

	cval, err := u.usecase.CreateSession(p.IDToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			ErrWithMessage{
				error: err,
				Why:   "",
			},
		)
	}

	sc := createCookie("exp-session", cval)
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
	userID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrWithMessage{
				error: err,
				Why:   "Parameter `userID` must be number",
			},
		)
	}

	// Fetch completion code by uid from DB
	code, err := u.usecase.GetCompletionCode(userID)
	if err != nil {
		// If given uid not found in DB
		if errors.Is(err, model.ErrNoSuchData) {
			return echo.NewHTTPError(http.StatusNotFound, ErrWithMessage{error: err, Why: "Not found"})
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{error: err, Why: "Request failed"},
		)
	}

	return c.JSON(http.StatusOK, code)
}
