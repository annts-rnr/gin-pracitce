package model

type Item struct {
	ID       int    `gorm:"primary_key;not null"`
	Title    string `gorm:"type:varchar(200);not null"`
	Contents string `gorm:"type:varchar(400)"`
	Price    int    `gorm:"type:integer"`
}
