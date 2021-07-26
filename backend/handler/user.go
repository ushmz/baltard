package handler

import (
	"database/sql"

	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/ymmt3-lab/koolhaas/backend/models"
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

// GroupCounts : Struct for group count
type GroupCounts struct {
	GroupId int `db:"group_id" json:"groupId"`
	Count   int `db:"count" json:"count"`
}

// allocateTask : Select one condition Id
func (h *Handler) allocateTask(c echo.Context) (int, error) {
	// tx : Begin transaction.
	tx := h.DB.MustBegin()

	// gc : Number of allocated users for each condition
	gc := GroupCounts{}
	// Fetch fewest task allocated number.
	err := tx.Get(&gc, `
		SELECT
			group_id,
			`+"`count`"+`
		FROM
			group_counts
		WHERE`+
		"   `count`"+` = (
				SELECT
					MIN(`+"`count`"+`)
				FROM
					group_counts
			)
		LIMIT
			1
	`)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		tx.Rollback()
		return 0, errors.New("Failed to allocate task")
	}

	// Increment condition count.
	_, err = tx.Exec(`
		UPDATE
			group_counts
		SET`+
		"   `count`"+` = ?
		WHERE
			group_id = ?
		`, gc.Count+1, gc.GroupId)
	if err != nil {
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		tx.Rollback()
		return 0, errors.New("Failed to allocate task")
	}

	// Commit changes in transaction.
	tx.Commit()

	return gc.GroupId, nil
}

// fetchTaskIdByCondition : Fetch task id that consist given group.
func (h *Handler) fetchTaskIdsByGroupId(groupId int) ([]int, error) {
	taskIds := []int{}
	err := h.DB.Select(&taskIds, `
		SELECT
			task_id
		FROM
			task_condition_relations
		WHERE
			group_id = ?
	`, groupId)
	if err != nil {
		return []int{0}, err
	}

	return taskIds, nil
}

// fetchTaskIdByCondition : Fetch task id that consist given group.
func (h *Handler) fetchConditionIdsByGroupId(groupId int) (int, error) {
	conditionId := []int{}
	err := h.DB.Select(&conditionId, `
		SELECT
			condition_id
		FROM
			task_condition_relations
		WHERE
			group_id = ?
	`, groupId)
	if err != nil {
		return 0, err
	}

	return conditionId[0], nil
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
	// eu : Exist user information
	eu := []models.ExistUser{}
	err := h.DB.Select(&eu, `
		SELECT
			id,
			uid,
			generated_secret
		FROM
			users
		WHERE
			uid = ?
	`, u.Uid)
	if err != nil {
		c.Echo().Logger.Errorf("Cannot detect user existence. : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if len(eu) == 0 {
		exist = false

		rand.Seed(time.Now().UnixNano())
		// randomNumber : Used as completion code
		randomNumber := rand.Intn(100000)
		// randomstr : Used as password (not necessary)
		randstr := generateRandomPasswd(12)

		// Insert user information
		rows, err := h.DB.Exec(`
		INSERT INTO
			users (
				uid,
				generated_secret
			)
		VALUES (?, ?)
		ON DUPLICATE
			KEY UPDATE
				generated_secret = ?`,
			u.Uid,
			randstr,
			randstr,
		)
		if err != nil {
			c.Echo().Logger.Errorf("Database Execution error : %v", err)
			return c.NoContent(http.StatusInternalServerError)
		}

		// Get last inserted id as user id
		insertedId, err := rows.LastInsertId()
		if err != nil {
			c.Echo().Logger.Errorf("Database Execution error : %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		// Insert completion code
		_, err = h.DB.Exec(`
		INSERT INTO 
			completion_codes (
				uid, 
				completion_code
			)
		VALUES (?, ?)`,
			u.Uid,
			randomNumber,
		)
		if err != nil {
			c.Echo().Logger.Errorf("Database Execution error : %v", err)
			return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
				Message: err.Error(),
			})
		}

		eu = append(eu, models.ExistUser{
			Id:     insertedId,
			Uid:    u.Uid,
			Secret: randstr,
		})
	}

	// groupId : Allocated group ID (consists of task IDs and condition ID)
	groupId, err := h.allocateTask(c)
	if err != nil {
		c.Echo().Logger.Errorf("Failed to allocate task  : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: err.Error(),
		})
	}
	// taskIds : Allocated task IDs
	taskIds, err := h.fetchTaskIdsByGroupId(groupId)
	if err != nil {
		c.Echo().Logger.Errorf("Failed to fetch task Ids : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: err.Error(),
		})
	}
	// conditionId : Allocated condition ID
	conditionId, err := h.fetchConditionIdsByGroupId(groupId)
	if err != nil {
		c.Echo().Logger.Errorf("Failed to fetch condition : %v", err)
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.UserResponse{
		Exist:       exist,
		UserId:      eu[0].Id,
		Secret:      eu[0].Secret,
		TaskIds:     taskIds,
		ConditionId: conditionId,
		GroupId:     groupId,
	})
}

// GetCompletionCode : Get users task completion code.
func (h *Handler) GetCompletionCode(c echo.Context) error {
	// id : Get id from path parameter
	id := c.Param("id")
	// idnum, _ := strconv.Atoi(id)

	var err error
	var completionCode models.CompletionCode

	// Fetch completion code by uid from DB
	query := `
	SELECT
		completion_code
	FROM
		completion_codes
	RIGHT JOIN
		users
	ON
		completion_codes.uid = users.uid
	WHERE
		users.id = ?`
	err = h.DB.Get(&completionCode, query, id)
	if err != nil {
		// If given uid not found in DB
		if err == sql.ErrNoRows {
			c.Echo().Logger.Infof("getCompletionCode uid %v not found", id)
			return c.NoContent(http.StatusNotFound)
		}
		c.Echo().Logger.Errorf("Database Execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	c.Echo().Logger.Info(completionCode)
	return c.JSON(http.StatusOK, completionCode)
}
