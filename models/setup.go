package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3307)/perpusgoo"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Buku{}, &Kategori{}, &User{})
	DB = db
}
