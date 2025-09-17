package repository

import (
	"CashMini/internal/user/models"
	"CashMini/internal/utils"
	"database/sql"
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
func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(utils.FIND_USERNAME_QUERY, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
