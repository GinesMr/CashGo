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
	repo := repository.NewUserRepository(db)    //Crud database methods
	service := services.NewUserService(repo)    //Business logic
	handler := handler2.NewUserHandler(service) //Request handler

	r.POST(utils.RouteRegister, handler.RegisterUser) //There is the UserEnpoints
}
