package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/intkuroky/blog-go/models"
	"golang.org/x/crypto/bcrypt"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})

	password, err := bcrypt.GenerateFromPassword([]byte("root"), bcrypt.DefaultCost)
	user := models.User{Name: "root", Password: string(password)}
	db.Create(&user)
}
