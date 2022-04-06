//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"math/rand"
	"strings"
	"time"

	"ratri/domain/authentication"
	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/pkg/errors"
)

type UserUsecase interface {
	FindByUid(uid string) (model.User, error)
	CreateUser(uid string) (model.User, error)
	CreateSession(idToken string) (string, error)
	AllocateTask() (model.TaskInfo, error)
	GetCompletionCode(userId int) (int, error)
}

type UserImpl struct {
	userRepository repo.UserRepository
	taskRepository repo.TaskRepository
	userAuth       authentication.UserAuthentication
}

func NewUserUsecase(
	userRepository repo.UserRepository,
	taskRepository repo.TaskRepository,
	userAuth authentication.UserAuthentication,
) UserUsecase {
	return &UserImpl{
		userRepository: userRepository,
		taskRepository: taskRepository,
		userAuth:       userAuth,
	}
}

// generateSecret : Generate password for new user.
func generateSecret(length, lower, upper, digits, symbols int) string {
	var (
		lowerCharSet = "abcdedfghijklmnopqrst"
		upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digitsSet    = "0123456789"
		symbolsSet   = "!@#$%&*"
		allCharSet   = lowerCharSet + upperCharSet + digitsSet + symbolsSet
	)

	var passwd strings.Builder

	for i := 0; i < lower; i++ {
		random := rand.Intn(len(lowerCharSet))
		passwd.WriteString(string(lowerCharSet[random]))
	}

	for i := 0; i < upper; i++ {
		random := rand.Intn(len(upperCharSet))
		passwd.WriteString(string(upperCharSet[random]))
	}

	for i := 0; i < digits; i++ {
		random := rand.Intn(len(digitsSet))
		passwd.WriteString(string(digitsSet[random]))
	}

	for i := 0; i < symbols; i++ {
		random := rand.Intn(len(symbolsSet))
		passwd.WriteString(string(symbolsSet[random]))
	}

	remaining := length - lower - upper - digits - symbols
	for i := 0; i < remaining; i++ {
		random := rand.Intn(len(allCharSet))
		passwd.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(passwd.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}

func (u *UserImpl) FindByUid(uid string) (model.User, error) {
	if u == nil {
		return model.User{}, errors.WithStack(model.ErrNilReciever)
	}

	user, err := u.userRepository.FindByUid(uid)
	if err != nil {
		return model.User{}, errors.WithStack(err)
	}

	return user, nil

}

func (u *UserImpl) CreateUser(uid string) (model.User, error) {
	if u == nil {
		return model.User{}, errors.WithStack(model.ErrNilReciever)
	}

	rand.Seed(time.Now().UnixNano())
	// randomNumber : Used as completion code
	randomNumber := rand.Intn(100000)
	// randomstr : Used as password
	secret := generateSecret(16, 2, 2, 2, 2)

	if err := u.userAuth.RegisterUser(uid, secret); err != nil {
		return model.User{}, err
	}

	user, err := u.userRepository.Create(uid)
	if err != nil {
		return model.User{}, err
	}

	// Insert completion code
	if u.userRepository.AddCompletionCode(user.Id, randomNumber); err != nil {
		return model.User{}, err
	}

	token, err := u.userAuth.GenerateToken(uid)
	if err != nil {
		return model.User{}, err
	}

	user.Token = token

	return user, nil
}

func (u *UserImpl) CreateSession(idToken string) (string, error) {
	if u == nil {
		return "", errors.WithStack(model.ErrNilReciever)
	}

	cookie, err := u.userAuth.GenerateSessionCookie(idToken, 1*time.Hour)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

func (u *UserImpl) AllocateTask() (model.TaskInfo, error) {
	if u == nil {
		return model.TaskInfo{}, errors.WithStack(model.ErrNilReciever)
	}

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
	if u == nil {
		return 0, errors.WithStack(model.ErrNilReciever)
	}

	return u.userRepository.GetCompletionCodeById(userId)
}
