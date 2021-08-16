package dao

import (
	"os"
	"testing"

	"baltard/database"

	_ "github.com/go-sql-driver/mysql"
)

var (
	logDao  Log
	serpDao Serp
	taskDao Task
	userDao User
)

func TestMain(m *testing.M) {
	cenv := database.NewMySQLConnectionEnv()
	db, err := cenv.ConnectDB()
	if err != nil {
		panic(err)
	}

	logDao = NewLog(db)
	serpDao = NewSerp(db)
	taskDao = NewTask(db)
	userDao = NewUser(db)

	os.Exit(m.Run())
}
