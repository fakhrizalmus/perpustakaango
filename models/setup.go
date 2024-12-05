package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3307)/uaspwd"))
	if err != nil {
		return
	}
	database.AutoMigrate(&Buku{}, &Kategori{})
	DB = database
}
