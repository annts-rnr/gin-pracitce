package repository

import (
	"fmt"

	"github.com/annts095/gin-practice/input"
	"github.com/annts095/gin-practice/model"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func (r *ItemRepository) FindById(input input.ItemIDInput) (*model.Item, error) {
	var item model.Item

	if err := r.DB.First(&item, input.ID).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) Save(item *model.Item) {
	r.DB.Create(item)
	fmt.Println("create success")
}
