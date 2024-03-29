package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func Userdb() {
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/go_data?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("not connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&User{})
}
