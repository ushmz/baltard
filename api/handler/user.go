package handler

import (
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"baltard/api/models"

	"github.com/labstack/echo"
)

// generateRandomPasswd : Generate random password its length equal to argument.
func generateRandomPasswd(l int) string {
	// Generate seed value.
	rand.Seed(time.Now().UnixNano())
	// letters : These letters are used for random password.
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// b : Generate random char in `l - 3` length.
	b := make([]rune, l-3)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	// To satisfy password policy, add some chars.
	return string(b) + "k2F"
}

// CreateUser : Register new user with crowd-sourcing service ID
func (h *Handler) CreateUser(c echo.Context) error {
	// u : Request body struct
	u := new(models.UserParam)
	// Bind request body parameters to struct
	if err := c.Bind(u); err != nil {
		c.Echo().Logger.Errorf("Error. Invalid request body. : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// exist : Given uid is already exist or not
	exist := true
	// Verbose
	var user models.ExistUser

	eu, err := h.User.FindByUid(u.Uid)
	if err != nil {
		if err != sql.ErrNoRows {
			c.Echo().Logger.Errorf("Cannot detect user existence. : %v", err)
			return c.NoContent(http.StatusInternalServerError)
		}
		exist = false

		rand.Seed(time.Now().UnixNano())
		// randomNumber : Used as completion code
		randomNumber := rand.Intn(100000)
		// randomstr : Used as password (not necessary)
		randstr := generateRandomPasswd(12)

		cu, err := h.User.Create(&models.User{Uid: u.Uid, Secret: randstr})
		if err != nil {
			c.Echo().Logger.Errorf("Database Execution error : %v", err)
			return c.NoContent(http.StatusInternalServerError)
		}

		// Insert completion code
		h.User.InsertCompletionCode(cu.Id, randomNumber)
		if err != nil {
			c.Echo().Logger.Errorf("Database Execution error : %v", err)
			return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
				Message: err.Error(),
			})
		}
		user = *cu
	} else {
		user = *eu
	}

	// groupId : Allocated group ID (consists of task IDs and condition ID)
	groupId, err := h.Task.AllocateTask()
	if err != nil {
		c.Echo().Logger.Errorf("Failed to allocate task  : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: err.Error(),
		})
	}
	// taskIds : Allocated task IDs
	taskIds, err := h.Task.FetchTaskIdsByGroupId(groupId)
	if err != nil {
		c.Echo().Logger.Errorf("Failed to fetch task Ids : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: err.Error(),
		})
	}
	// conditionId : Allocated condition ID
	conditionId, err := h.Condition.FetchConditionIdByGroupId(groupId)
	if err != nil {
		c.Echo().Logger.Errorf("Failed to fetch condition : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.UserResponse{
		Exist:       exist,
		UserId:      user.Id,
		Secret:      user.Secret,
		TaskIds:     taskIds,
		ConditionId: conditionId,
		GroupId:     groupId,
	})
}

// GetCompletionCode : Get users task completion code.
func (h *Handler) GetCompletionCode(c echo.Context) error {
	// id : Get id from path parameter
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "Parameter `userId` must be number",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	// Fetch completion code by uid from DB
	code, err := h.User.GetCompletionCodeById(userId)
	if err != nil {
		// If given uid not found in DB
		if err == sql.ErrNoRows {
			c.Echo().Logger.Infof("uid %v not found", id)
			return c.NoContent(http.StatusNotFound)
		}
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, code)
}
