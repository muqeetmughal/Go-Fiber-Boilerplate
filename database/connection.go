package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(){
	_ , err := gorm.Open(mysql.Open("root:12345678@/gofiberbp"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}