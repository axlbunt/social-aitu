package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"social/models"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/social_db"), &gorm.Config{})

	if err != nil {
		panic("No connection to DB!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
