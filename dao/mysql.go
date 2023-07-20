package dao

import (
	"VisitorsManagementSystem/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "callersystem"
)

func InitDB() *gorm.DB {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/callersystem?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Visitor{}, models.Admin{})
	return db
}
