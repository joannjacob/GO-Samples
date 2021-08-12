package db

import (
	"database/sql"
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func CreateDatabase() (*sql.DB, error) {
	// serverName := "localhost"
	// port ="5432"
	user := "postgres"
	password := "postgres"
	dbName := "chi_mux_sample"

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)

	connectionString := fmt.Sprintf(" user=%s "+
		"password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil

}
