package mysql

import (
	"database/sql"

	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// UserRepositoryImpl : Struct for use repository
type UserRepositoryImpl struct {
	DB *sqlx.DB
}

// NewUserRepository : Return new UserRepository
func NewUserRepository(db *sqlx.DB) repo.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

// Create : Create new user
func (u *UserRepositoryImpl) Create(uid string) (model.User, error) {
	user := model.User{}
	// [TODO] Save completion code at the same time
	rows, err := u.DB.Exec(`INSERT INTO users (uid) VALUES (?) `, uid)
	if err != nil {
		return user, err
	}

	insertedID, err := rows.LastInsertId()
	if err != nil {
		return user, err
	}

	user = model.User{
		ID:  int(insertedID),
		UID: uid,
	}
	return user, nil
}

// FindByID : Find an user by ID
func (u *UserRepositoryImpl) FindByID(userID int) (model.User, error) {
	user := model.User{}
	row := u.DB.QueryRowx(`SELECT id, uid FROM users WHERE id = ?`, userID)
	if err := row.StructScan(&user); err != nil {
		if err == sql.ErrNoRows {
			return user, model.ErrNoSuchData
		}
		return user, err
	}
	return user, nil
}

// FindByUID : Find an user by UID
func (u *UserRepositoryImpl) FindByUID(uid string) (model.User, error) {
	user := model.User{}
	err := u.DB.Get(&user, `SELECT id, uid FROM users WHERE uid = ?`, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, model.ErrNoSuchData
		}
		return user, err
	}
	return user, nil
}

// AddCompletionCode : Add completion code
func (u *UserRepositoryImpl) AddCompletionCode(userID, code int) error {
	_, err := u.DB.Exec(`
		INSERT INTO 
			completion_codes (
				user_id, 
				completion_code
			)
		VALUES (?, ?)`,
		userID,
		code,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetCompletionCodeByID : Get task completion code by user ID
func (u *UserRepositoryImpl) GetCompletionCodeByID(userID int) (int, error) {
	var code sql.NullInt64
	row := u.DB.QueryRow(`
		SELECT
			completion_code
		FROM
			completion_codes
		WHERE
			user_id = ?
	`, userID)

	if err := row.Scan(&code); err != nil {
		if err == sql.ErrNoRows {
			return 0, model.ErrNoSuchData
		}
		return 0, errors.WithStack(err)
	}

	if !code.Valid {
		return 42, errors.WithStack(model.ErrInternalServerError)
	}
	return int(code.Int64), nil
}
