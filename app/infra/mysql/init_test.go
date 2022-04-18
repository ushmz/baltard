package mysql_test

import (
	"os"
	"testing"

	repo "ratri/domain/repository"
	"ratri/infra/mysql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	linkedPageDao repo.LinkedPageRepository
	logDao        repo.LogRepository
	serpDao       repo.SerpRepository
	taskDao       repo.TaskRepository
	userDao       repo.UserRepository
)

func TestMain(m *testing.M) {
	db, err := mysql.New()
	if err != nil {
		panic(err)
	}

	linkedPageDao = mysql.NewLinkedPageRepository(db)
	logDao = mysql.NewLogRepository(db)
	serpDao = mysql.NewSerpRepository(db)
	taskDao = mysql.NewTaskRepository(db)
	userDao = mysql.NewUserRepository(db)

	os.Exit(m.Run())
}
