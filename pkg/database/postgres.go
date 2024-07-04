package database

import (
	"database/sql"
	"fmt"
	"strconv"
)

func OpenDatabase() (*sql.DB, error) {
	port, err := strconv.Atoi("5432")
	if err != nil {
		return nil, err
	}

	return sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			"localhost", port, "postgres", "admin123", "lms_db",
		),
	)
}
