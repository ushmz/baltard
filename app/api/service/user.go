package service

import (
	"baltard/api/dao"
	"baltard/api/model"

	"database/sql"
	"math/rand"
	"time"
)

type User interface {
	GenerateRandomPasswd(l int) string
	FindByUid(uid string) (*model.User, bool, error)
	CreateUser(uid string) (*model.User, error)
	AllocateTask() (model.TaskInfo, error)
	GetCompletionCode(userId int) (int, error)
}

type UserImpl struct {
	userDao dao.User
	taskDao dao.Task
}

func NewUserService(userDao dao.User, taskDao dao.Task) User {
	return &UserImpl{userDao: userDao, taskDao: taskDao}
}

// generateRandomPasswd : Generate random password its length equal to argument.
func (u *UserImpl) GenerateRandomPasswd(l int) string {
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

func (u *UserImpl) FindByUid(uid string) (*model.User, bool, error) {
	user, err := u.userDao.FindByUid(uid)
	if err != nil {
		if err != sql.ErrNoRows {
			return &model.User{}, false, err
		}
		return &model.User{}, false, nil
	}

	return user, true, nil

}

func (u *UserImpl) CreateUser(uid string) (*model.User, error) {
	rand.Seed(time.Now().UnixNano())
	// randomNumber : Used as completion code
	randomNumber := rand.Intn(100000)
	// randomstr : Used as password (not necessary)
	randstr := u.GenerateRandomPasswd(12)

	cu, err := u.userDao.Create(uid, randstr)
	if err != nil {
		return &model.User{}, err
	}

	// Insert completion code
	u.userDao.InsertCompletionCode(cu.Id, randomNumber)
	if err != nil {
		return &model.User{}, err
	}
	return cu, nil
}

func (u *UserImpl) AllocateTask() (model.TaskInfo, error) {
	// groupId : Allocated group ID (consists of task IDs and condition ID)
	groupId, err := u.taskDao.AllocateTask()
	if err != nil {
		return model.TaskInfo{}, err
	}

	// taskIds : Allocated task IDs
	taskIds, err := u.taskDao.FetchTaskIdsByGroupId(groupId)
	if err != nil {
		return model.TaskInfo{}, err
	}

	// conditionId : Allocated condition ID
	conditionId, err := u.taskDao.FetchConditionIdByGroupId(groupId)
	if err != nil {
		return model.TaskInfo{}, err
	}

	return model.TaskInfo{ConditionId: conditionId, GroupId: groupId, TaskIds: taskIds}, nil
}

func (u *UserImpl) GetCompletionCode(userId int) (int, error) {
	return u.userDao.GetCompletionCodeById(userId)
}
