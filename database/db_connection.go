package database

import (
	"fmt"
	logger    "../customlogger"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var logs = logger.GetInstance("SYSTEMS -")

const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "user"
    DB_USER = "root"
    DB_PASS = ""
)

/*Create mysql connection*/
func CreateCon() *gorm.DB{
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME))
	if err != nil {
		logs.Println(err)
	}
	return db
}
