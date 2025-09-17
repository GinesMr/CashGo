package user

import (
	handler2 "CashMini/internal/user/handler"
	"CashMini/internal/user/repository"
	"CashMini/internal/user/services"
	"CashMini/internal/utils"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	repo := repository.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handler2.NewUserHandler(service)

	r.POST(utils.RouteRegister, handler.RegisterUser)
}
