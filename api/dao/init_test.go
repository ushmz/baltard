package dao

import (
	"os"
	"testing"

	"baltard/database"

	_ "github.com/go-sql-driver/mysql"
)

var (
	answerDao    Answer
	conditionDao Condition
	logDao       Log
	serpDao      Serp
	taskDao      Task
	userDao      User
)

func TestMain(m *testing.M) {
	cenv := database.NewMySQLConnectionEnv()
	db, err := cenv.ConnectDB()
	if err != nil {
		panic(err)
	}

	answerDao = NewAnswer(db)
	conditionDao = NewCondition(db)
	logDao = NewLog(db)
	serpDao = NewSerp(db)
	taskDao = NewTask(db)
	userDao = NewUser(db)

	os.Exit(m.Run())
}
