//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"math/rand"
	"time"

	"ratri/domain/model"
	repo "ratri/domain/repository"
)

type UserUsecase interface {
	GenerateRandomPasswd(l int) string
	FindByUid(uid string) (model.User, bool, error)
	CreateUser(uid string) (model.User, error)
	AllocateTask() (model.TaskInfo, error)
	GetCompletionCode(userId int) (int, error)
}

type UserImpl struct {
	userRepository repo.UserRepository
	taskRepository repo.TaskRepository
}

func NewUserUsecase(userRepository repo.UserRepository, taskRepository repo.TaskRepository) UserUsecase {
	return &UserImpl{userRepository: userRepository, taskRepository: taskRepository}
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

func (u *UserImpl) FindByUid(uid string) (model.User, bool, error) {
	user, err := u.userRepository.FindByUid(uid)
	if err != nil {
		return model.User{}, false, nil
	}

	return user, true, nil

}

func (u *UserImpl) CreateUser(uid string) (model.User, error) {
	zv := model.User{}

	rand.Seed(time.Now().UnixNano())
	// randomNumber : Used as completion code
	randomNumber := rand.Intn(100000)
	// randomstr : Used as password (not necessary)
	randstr := u.GenerateRandomPasswd(12)

	cu, err := u.userRepository.Create(uid, randstr)
	if err != nil {
		return zv, err
	}

	// Insert completion code
	u.userRepository.AddCompletionCode(cu.Id, randomNumber)
	if err != nil {
		return zv, err
	}
	return cu, nil
}

func (u *UserImpl) AllocateTask() (model.TaskInfo, error) {
	// groupId : Allocated group ID (consists of task IDs and condition ID)
	groupId, err := u.taskRepository.UpdateTaskAllocation()
	if err != nil {
		return model.TaskInfo{}, err
	}

	// taskIds : Allocated task IDs
	taskIds, err := u.taskRepository.GetTaskIdsByGroupId(groupId)
	if err != nil {
		return model.TaskInfo{}, err
	}

	// conditionId : Allocated condition ID
	conditionId, err := u.taskRepository.GetConditionIdByGroupId(groupId)
	if err != nil {
		return model.TaskInfo{}, err
	}

	return model.TaskInfo{ConditionId: conditionId, GroupId: groupId, TaskIds: taskIds}, nil
}

func (u *UserImpl) GetCompletionCode(userId int) (int, error) {
	return u.userRepository.GetCompletionCodeById(userId)
}
