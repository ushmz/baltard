package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/ymmt3-lab/koolhaas/backend/api/models"
)

type User interface {
	Create(user *models.User) (*models.ExistUser, error)
	FindById(UserId int) (*models.ExistUser, error)
	FindByUid(uid string) (*models.ExistUser, error)
	InsertCompletionCode(userId, code int) error
	GetCompletionCodeById(userId int) (int, error)
}

type UserImpl struct {
	DB *sqlx.DB
}

func NewUser(db *sqlx.DB) User {
	return &UserImpl{DB: db}
}

func (u UserImpl) Create(user *models.User) (*models.ExistUser, error) {
	rows, err := u.DB.Exec(`
		INSERT INTO
			users (
				uid,
				generated_secret
			)
		VALUES (?, ?)
	`,
		user.Uid,
		user.Secret,
	)
	if err != nil {
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	eu := models.ExistUser{
		Id:     int(insertedId),
		Uid:    user.Uid,
		Secret: user.Secret,
	}
	return &eu, nil
}

func (u UserImpl) FindById(userId int) (*models.ExistUser, error) {
	// [TODO] `ExistUser` might be verbose struct
	user := models.ExistUser{}
	row := u.DB.QueryRow(`
		SELECT
			id,
			uid,
			generated_secret
		FROM
			users
		WHERE
			id = ?
	`, userId)
	if err := row.Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserImpl) FindByUid(uid string) (*models.ExistUser, error) {
	// [TODO] `ExistUser` might be verbose struct
	user := models.ExistUser{}
	row := u.DB.QueryRow(`
		SELECT
			id,
			uid,
			generated_secret
		FROM
			users
		WHERE
			uid = ?
	`, uid)
	if err := row.Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserImpl) InsertCompletionCode(userId, code int) error {
	_, err := u.DB.Exec(`
		INSERT INTO 
			completion_codes (
				uid, 
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

func (u UserImpl) GetCompletionCodeById(userId int) (int, error) {
	var code int
	row := u.DB.QueryRow(`
		SELECT
			completion_code
		FROM
			completion_codes
		RIGHT JOIN
			users
		ON
			completion_codes.uid = users.uid
		WHERE
			users.id = ?
	`, userId)

	if err := row.Scan(&code); err != nil {
		return 0, err
	}
	return code, nil
}
