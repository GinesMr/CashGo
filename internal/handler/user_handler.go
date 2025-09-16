package handler

import (
	"CashMini/internal/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var request dtos.RegisterRequest
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": "IAM GAY"})
	/*
		user, err := h.service.RegisterUser(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}*/

	/*	c.JSON(http.StatusCreated, gin.H{"id": user.ID, "username": user.Username})*/
}
