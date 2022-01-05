package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	ID       int    `gorm:"primary_key;not null"`
	Title    string `gorm:"type:varchar(200);not null"`
	Contents string `gorm:"type:varchar(400)"`
	Price    string `gorm:"type:integer"`
}

func getGormConnect() *gorm.DB {
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
	db := getGormConnect()
	db.AutoMigrate(&Item{})

	fmt.Println("migrate success")
}
