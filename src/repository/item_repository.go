package repository

import (
	"fmt"

	"github.com/annts095/gin-practice/model"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func (r *ItemRepository) Save(item *model.Item) {
	r.DB.Create(item)
	fmt.Println("create success")
}
