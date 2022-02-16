package utils

import (
	"fmt"
	"gorm.io/gorm"
	"helpstack/models"
	"helpstack/pkg/article"
)

func MigrateDb(Db *gorm.DB){
	err := Db.AutoMigrate(&models.Role{})
	if err != nil {
		fmt.Println("failed to automigrate role model:", err.Error())
		return
	}
	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("failed to automigrate user model:", err.Error())
		return
	}
	err = Db.AutoMigrate(&article.Article{})
	if err != nil {
		fmt.Println("failed to automigrate user model:", err.Error())
		return
	}
}