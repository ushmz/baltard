package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"baltard/api/model"
	"baltard/api/service"

	"github.com/labstack/echo"
)

type User struct {
	service service.User
}

func NewUserHandler(userService service.User) *User {
	return &User{service: userService}
}

// CreateUser : Register new user with crowd-sourcing service ID
func (u *User) CreateUser(c echo.Context) error {
	// u : Request body struct
	param := new(model.UserParam)
	// Bind request body parameters to struct
	if err := c.Bind(param); err != nil {
		c.Echo().Logger.Errorf("Error. Invalid request body. : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// Verbose
	var user model.User

	// exist : Given uid is already exist or not
	eu, exist, err := u.service.FindByUid(param.Uid)
	if err != nil {
		c.Echo().Logger.Errorf("Failed to detect user existence : %v", err)
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: err.Error(),
		})
	}

	if exist {
		user = *eu
	} else {
		nu, err := u.service.CreateUser(param.Uid)
		if err != nil {
			c.Echo().Logger.Errorf("Failed to create new user : %v", err)
			return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
				Message: err.Error(),
			})
		}
		user = *nu
	}

	info, err := u.service.AllocateTask()
	if err != nil {
		c.Echo().Logger.Errorf("Failed to allocate task  : %v", err)
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: err.Error(),
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
	code, err := u.service.GetCompletionCode(userId)
	if err != nil {
		// If given uid not found in DB
		if err == sql.ErrNoRows {
			return c.NoContent(http.StatusNotFound)
		}
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, code)
}
