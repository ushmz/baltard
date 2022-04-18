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

// UserUsecase : Abstract operations that user usecase should have
type UserUsecase interface {
	FindByUID(uid string) (model.User, error)
	CreateUser(uid string) (model.User, error)
	CreateSession(idToken string) (string, error)
	AllocateTask() (model.TaskInfo, error)
	GetCompletionCode(userID int) (int, error)
}

// UserImpl : Implemention of user usecase
type UserImpl struct {
	userRepository repo.UserRepository
	taskRepository repo.TaskRepository
	userAuth       authentication.UserAuthentication
}

// NewUserUsecase : Return new user usecase struct
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

// FindByUID : Get a user by UID
func (u *UserImpl) FindByUID(uid string) (model.User, error) {
	if u == nil {
		return model.User{}, errors.WithStack(model.ErrNilReceiver)
	}

	user, err := u.userRepository.FindByUID(uid)
	if err != nil {
		return model.User{}, errors.WithStack(err)
	}

	return user, nil

}

// CreateUser : Create new user on this system
func (u *UserImpl) CreateUser(uid string) (model.User, error) {
	if u == nil {
		return model.User{}, errors.WithStack(model.ErrNilReceiver)
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
	if u.userRepository.AddCompletionCode(user.ID, randomNumber); err != nil {
		return model.User{}, err
	}

	token, err := u.userAuth.GenerateToken(uid)
	if err != nil {
		return model.User{}, err
	}

	user.Token = token

	return user, nil
}

// CreateSession : Create session cookie
func (u *UserImpl) CreateSession(idToken string) (string, error) {
	if u == nil {
		return "", errors.WithStack(model.ErrNilReceiver)
	}

	cookie, err := u.userAuth.GenerateSessionCookie(idToken, 1*time.Hour)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

// AllocateTask : Allocate tasks to user
func (u *UserImpl) AllocateTask() (model.TaskInfo, error) {
	if u == nil {
		return model.TaskInfo{}, errors.WithStack(model.ErrNilReceiver)
	}

	// groupID : Allocated group ID (consists of task IDs and condition ID)
	groupID, err := u.taskRepository.UpdateTaskAllocation()
	if err != nil {
		return model.TaskInfo{}, err
	}

	// taskIDs : Allocated task IDs
	taskIDs, err := u.taskRepository.GetTaskIDsByGroupID(groupID)
	if err != nil {
		return model.TaskInfo{}, err
	}

	// conditionID : Allocated condition ID
	conditionID, err := u.taskRepository.GetConditionIDByGroupID(groupID)
	if err != nil {
		return model.TaskInfo{}, err
	}

	return model.TaskInfo{ConditionID: conditionID, GroupID: groupID, TaskIDs: taskIDs}, nil
}

// GetCompletionCode : Get user task completion code
func (u *UserImpl) GetCompletionCode(userID int) (int, error) {
	if u == nil {
		return 0, errors.WithStack(model.ErrNilReceiver)
	}

	return u.userRepository.GetCompletionCodeByID(userID)
}
