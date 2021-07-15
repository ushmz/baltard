package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

var db *sqlx.DB
var mySQLConnectionData *MySQLConnectionEnv

var (
	migrations = &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
)

type MySQLConnectionEnv struct {
	Host           string
	Port           string
	User           string
	DBName         string
	Password       string
	ConnectionName string
}


func NewMySQLConnectionEnv() *MySQLConnectionEnv {
	return &MySQLConnectionEnv{
		Host:           getEnv("MYSQL_HOST", "0.0.0.0"),
		Port:           getEnv("MYSQL_PORT", "3366"),
		User:           getEnv("MYSQL_USER", "koolhaas"),
		DBName:         getEnv("MYSQL_DBNAME", "koolhaas"),
		Password:       getEnv("MYSQL_PASS", "koolhaas"),
		ConnectionName: getEnv("CONNECTION_NAME", "default"),
	}
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultValue
}

func (mc *MySQLConnectionEnv) ConnectDB() (*sqlx.DB, error) {
	var dsn string
	if getEnv("ENV", "") == "prd" {
		dsn = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True", mc.User, mc.Password, mc.ConnectionName, mc.DBName)
		fmt.Printf(dsn)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mc.User, mc.Password, mc.Host, mc.Port, mc.DBName)
	}
	return sqlx.Open("mysql", dsn)
}

func New() *sqlx.DB {
	mySQLConnectionData = NewMySQLConnectionEnv()

	var err error
	db, err = mySQLConnectionData.ConnectDB()
	if err != nil {
		log.Printf("DB connection failed: %v", err)
	}

	if getEnv("ENV", "") == "prd" {
		appliedCount, err := migrate.Exec(db.DB, "mysql", migrations, migrate.Up)
		if err != nil {
			log.Printf("DB migration failed: %v", err)
		}
		log.Printf("Applied %v migrations", appliedCount)
	}

	// DB指定
	db.Exec("USE koolhaas")

	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1500)
	return db
}