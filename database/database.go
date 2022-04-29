package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(){
	dsn := "host=localhost user=postgres password=yutakaki dbname=findMentor2 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println(err)
		return
	}
	Db = db
}