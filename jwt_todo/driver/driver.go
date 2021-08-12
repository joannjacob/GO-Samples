package driver

import (
	"fmt"
	models "jwt_todo/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// DB ...
type DB struct {
	SQL *gorm.DB
}

// DBConn ...
var dbConn = &DB{}

var db *gorm.DB
var err error

// ConnectSQL ...
func ConnectSQL(host, port, user, pass, dbname string) (*gorm.DB, error) {

	db, err = gorm.Open(
		"postgres",
		"host="+host+" user="+user+
			" dbname="+dbname+" sslmode=disable password="+
			pass)

	if err != nil {
		fmt.Println("DB connection failed")
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(models.Todo{}, models.User{}, models.TokenDetails{}, models.AccessDetails{})
	// dbConn.SQL = db
	return db, err
}
