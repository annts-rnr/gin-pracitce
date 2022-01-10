package controller

import (
	"net/http"
	"strconv"

	"github.com/annts095/gin-practice/input"
	"github.com/annts095/gin-practice/model"
	"github.com/annts095/gin-practice/repository"
	"github.com/gin-gonic/gin"
)

type ItemController struct {
	Context    *gin.Context
	Repository *repository.ItemRepository
}
}

func (c *ItemController) Create() {
	var input input.ItemCreateInput

	if err := c.Context.ShouldBindJSON(&input); err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}
	item := model.Item{
		Title:    input.Item.Title,
		Contents: input.Item.Contents,
		Price:    input.Item.Price,
	}

	c.Repository.Save(&item)
	c.Context.JSON(http.StatusOK, gin.H{
		"id":       item.ID,
		"title":    item.Title,
		"contents": item.Contents,
		"price":    item.Price,
	})
}
