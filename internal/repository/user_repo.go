package repository

import (
	"CashMini/internal/models"
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
	).Scan(&user.ID)
}
