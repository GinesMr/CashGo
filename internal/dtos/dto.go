package dtos

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"hasher" binding:"required"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
