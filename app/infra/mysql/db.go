package mysql

import (
	"fmt"
	"ratri/config"
	"ratri/domain/model"
	"time"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	migrations = &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
)

// ConnectDB : Return MySQL connection object with a timeout (30 sec.)
func connectDB() (*sqlx.DB, error) {
	conf := config.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetString("db.user"),
		conf.GetString("db.password"),
		conf.GetString("db.host"),
		conf.GetString("db.port"),
		conf.GetString("db.database"),
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err != nil {
			fmt.Println("DB is not ready. Retry connecting...")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println("Success to connect DB")
		return db, nil
	}

	return nil, model.ErrDBConnectionFailed
}

// New : Apply migration if it runs in production stage, then return DB connection
func New() (*sqlx.DB, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	conf := config.GetConfig()
	if conf.GetString("env") == "prod" {
		appliedCount, err := migrate.Exec(db.DB, "mysql", migrations, migrate.Up)
		if err != nil {
			return nil, fmt.Errorf("DB migration failed: %v", err)
		}
		fmt.Printf("Applied %v migrations", appliedCount)
	}

	// db.SetMaxOpenConns(25500)
	// db.SetMaxIdleConns(25500)
	return db, nil
}
