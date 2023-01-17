package mq

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestMysql(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:lch123456@tcp(172.30.163.23:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Community{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.UserBasic{})

}
