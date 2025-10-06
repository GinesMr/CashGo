package repository

import (
	"CashMini/internal/user/models"
	"CashMini/internal/utils"
	"database/sql"
	"fmt"
)

// /CRUD
type UserRepository struct {
	db *sql.DB ///There is we create the instace of the database
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.UserModel) error {
	return r.db.QueryRow(utils.CREATE_USER_QUERY,
		user.Username,
		user.Password,
		user.Email,
	).Scan(&user.ID)
}
func (r *UserRepository) DeleteUserByUsernamePassword(username string, password string) (bool, error) {
	err := r.db.QueryRow(utils.DELETE_USER_QUERY).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *UserRepository) findIdByUsernamePassword(username string, password string) (int64, error) {
	var idUser int64
	err := r.db.QueryRow(utils.FIND_ID_BY_USERNAME_PASSWORD_QUERY, username, password).Scan(&idUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("user not found or invalid password")
		}
		return 0, err
	}
	return idUser, nil
}
func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(utils.FIND_USERNAME_QUERY, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
