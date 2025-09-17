package utils

const (
	CREATE_USER_QUERY   = "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id;"
	FIND_USERNAME_QUERY = "SELECT EXISTS ( SELECT 1 FROM users WHERE username = $1);"
)
