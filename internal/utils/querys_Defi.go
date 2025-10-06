package utils

const (
	CREATE_USER_QUERY                  = "INSERT INTO users (username,password_hash,email) VALUES ($1, $2, $3) RETURNING id;"
	DELETE_USER_QUERY                  = "DELETE FROM users WHERE  = $1 RETURNING id;"
	FIND_USERNAME_QUERY                = "SELECT EXISTS ( SELECT 1 FROM users WHERE username = $1, password = $2 );"
	FIND_ID_BY_USERNAME_PASSWORD_QUERY = "SELECT id FROM users WHERE username = $1, password = $2 LIMIT 1;"
)
