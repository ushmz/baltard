package mysql_test

import (
	"os"
	"testing"

	repo "ratri/src/domain/repository"
	"ratri/src/infra/mysql"

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
	cenv := mysql.NewMySQLConnectionEnv()
	db, err := cenv.ConnectDB()
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
