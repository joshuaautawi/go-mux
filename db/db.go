package db

import (
	"github.com/joshuaautawi/warehouse/cmd/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/warehouse"))
	if err != nil{
		panic(err)
	}
	db.AutoMigrate(models.Product{})
	DB = db
}