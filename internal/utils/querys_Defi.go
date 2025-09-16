package utils

const (
	CREATE_USER_QUERY = "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id;"
)
