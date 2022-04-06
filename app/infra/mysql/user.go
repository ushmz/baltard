package mysql

import (
	"database/sql"

	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repo.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) Create(uid string) (model.User, error) {
	user := model.User{}
	// [TODO] Save completion code at the same time
	rows, err := u.DB.Exec(`INSERT INTO users (uid) VALUES (?) `, uid)
	if err != nil {
		return user, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		return user, err
	}

	user = model.User{
		Id:  int(insertedId),
		Uid: uid,
	}
	return user, nil
}

func (u *UserRepositoryImpl) FindById(userId int) (model.User, error) {
	user := model.User{}
	row := u.DB.QueryRowx(`SELECT id, uid FROM users WHERE id = ?`, userId)
	if err := row.StructScan(&user); err != nil {
		if err == sql.ErrNoRows {
			return user, model.ErrNoSuchData
		}
		return user, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) FindByUid(uid string) (model.User, error) {
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

func (u *UserRepositoryImpl) AddCompletionCode(userId, code int) error {
	_, err := u.DB.Exec(`
		INSERT INTO 
			completion_codes (
				user_id, 
				completion_code
			)
		VALUES (?, ?)`,
		userId,
		code,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) GetCompletionCodeById(userId int) (int, error) {
	var code sql.NullInt64
	row := u.DB.QueryRow(`
		SELECT
			completion_code
		FROM
			completion_codes
		WHERE
			user_id = ?
	`, userId)

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
