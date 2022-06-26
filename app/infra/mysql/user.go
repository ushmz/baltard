package mysql

import (
	"database/sql"
	"fmt"

	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
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
	if u == nil {
		return user, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

	// [TODO] Save completion code at the same time
	rows, err := u.DB.Exec(`INSERT INTO users (uid) VALUES (?) `, uid)
	if err != nil {
		return user, fmt.Errorf("Try to insert new user row: %w", err)
	}

	insertedID, err := rows.LastInsertId()
	if err != nil {
		return user, fmt.Errorf("Try to get last inserted ID: %w", err)
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
	if u == nil {
		return user, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

	row := u.DB.QueryRowx(`SELECT id, uid FROM users WHERE id = ?`, userID)
	if err := row.StructScan(&user); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("User with ID (%d) is not found: %w", userID, model.ErrNoSuchData)
		}
		return user, fmt.Errorf("Try to get user with ID (%d): %w", userID, err)
	}
	return user, nil
}

// FindByUID : Find an user by UID
func (u *UserRepositoryImpl) FindByUID(uid string) (model.User, error) {
	user := model.User{}
	if u == nil {
		return user, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

	if err := u.DB.Get(&user, `SELECT id, uid FROM users WHERE uid = ?`, uid); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("User with UID (%s) is not found: %w", uid, model.ErrNoSuchData)
		}
		return user, fmt.Errorf("Try to get user with UID (%s): %w", uid, err)
	}
	return user, nil
}

// AddCompletionCode : Add completion code
func (u *UserRepositoryImpl) AddCompletionCode(userID, code int) error {
	if u == nil {
		return fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
		return fmt.Errorf("Try to insert completion code (%d) to User (%d): %w", code, userID, err)
	}
	return nil
}

// GetCompletionCodeByID : Get task completion code by user ID
func (u *UserRepositoryImpl) GetCompletionCodeByID(userID int) (int, error) {
	if u == nil {
		return 0, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
			return 0, fmt.Errorf("Completion code for user(%d) is not found: %w", userID, model.ErrNoSuchData)
		}
		return 0, fmt.Errorf("Try to get completion code for user (%d): %w", userID, err)
	}

	if !code.Valid {
		return 42, fmt.Errorf("Invalid completion code for user(%d): %w", userID, model.ErrInternal)
	}

	return int(code.Int64), nil
}
