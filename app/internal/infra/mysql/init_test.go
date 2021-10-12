package mysql

import (
	"os"
	"testing"

	repo "ratri/internal/domain/repository"

	_ "github.com/go-sql-driver/mysql"
)

var (
	logDao  repo.LogRepository
	serpDao repo.SerpRepository
	taskDao repo.TaskRepository
	userDao repo.UserRepository
)

func TestMain(m *testing.M) {
	cenv := NewMySQLConnectionEnv()
	db, err := cenv.ConnectDB()
	if err != nil {
		panic(err)
	}

	logDao = NewLogRepository(db)
	serpDao = NewSerpRepository(db)
	taskDao = NewTaskRepository(db)
	userDao = NewUserRepository(db)

	os.Exit(m.Run())
}
