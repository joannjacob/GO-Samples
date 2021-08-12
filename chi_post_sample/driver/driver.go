package driver

import (
	// "chi_post_sample/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DB ...
type DB struct {
	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}

var db *sql.DB
var err error

// ConnectSQL ...
func ConnectSQL(host, port, user, pass, dbname string) (*sql.DB, error) {

	db, err = sql.Open(
		"postgres",
		"host="+host+" user="+user+
			" dbname="+dbname+" sslmode=disable password="+
			pass)

	if err != nil {
		fmt.Println("DB connection failed")
		panic("failed to connect database")
	}

	defer db.Close()

	// db.AutoMigrate(models.Post{})
	// dbConn.SQL = db
	return db, err
}
