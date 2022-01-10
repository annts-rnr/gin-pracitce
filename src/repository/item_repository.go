package repository

import (
	"github.com/annts095/gin-practice/model"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func (r *ItemRepository) FindById(id int) (*model.Item, error) {
	var item model.Item

	if err := r.DB.First(&item, id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) Save(item *model.Item) {
	r.DB.Save(item)
}

func (r *ItemRepository) Delete(item *model.Item) error {
	if err := r.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}
