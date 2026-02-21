package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rhodamineb13/backend-test/utils"
)

var DB *gorm.DB

func ConnectSQL() {
	url := utils.MySQLURL
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		panic(err)
	}

	DB = db
}
