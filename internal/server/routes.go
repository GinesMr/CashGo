package server

import (
	"CashMini/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.db.Health())
	})
	user.RegisterRoutes(r, s.db.DB())
	return r
}
