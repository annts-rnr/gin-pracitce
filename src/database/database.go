package database

import (
	"fmt"
	"os"

	"github.com/annts095/gin-practice/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetGormConnect() *gorm.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PORT := os.Getenv("DB_PORT")
	HOST := os.Getenv("DB_HOST")
	PROTOCOL := "tcp(" + HOST + ":" + PORT + ")"
	DBNAME := "database"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("db connected: ", &db)
	return db
}

func ItemMigrate() {
	db := GetGormConnect()
	db.AutoMigrate(&model.Item{})

	fmt.Println("migrate success")
}
