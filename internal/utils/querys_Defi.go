package utils

const (
	CREATE_USER_QUERY   = "INSERT INTO users (username,password_hash,email) VALUES ($1, $2, $3) RETURNING id;"
	FIND_USERNAME_QUERY = "SELECT EXISTS ( SELECT 1 FROM users WHERE username = $1);"
)
