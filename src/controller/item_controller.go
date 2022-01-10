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

func (c *ItemController) Show() {
	var input input.ItemIDInput

	if err := c.Context.ShouldBindUri(&input); err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	item, err := c.Repository.FindById(input.ID)
	if err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	c.Context.JSON(http.StatusOK, gin.H{
		"id":       item.ID,
		"title":    item.Title,
		"contents": item.Contents,
		"price":    item.Price,
	})
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

func (c *ItemController) Update() {
	var input input.ItemUpdateInput

	if err := c.Context.ShouldBindJSON(&input); err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	if err := c.Context.ShouldBindUri(&input); err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	item, err := c.Repository.FindById(input.ID)
	if err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	item.Title = input.Item.Contents
	item.Contents = input.Item.Contents
	item.Price = input.Item.Price

	c.Repository.Save(item)
	c.Context.JSON(http.StatusOK, gin.H{
		"id":       item.ID,
		"title":    item.Title,
		"contents": item.Contents,
		"price":    item.Price,
	})
}

func (c *ItemController) Delete() {
	var input input.ItemIDInput

	if err := c.Context.ShouldBindUri(&input); err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	item, err := c.Repository.FindById(input.ID)
	if err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	if err := c.Repository.Delete(item); err != nil {
		h := map[string]string{
			"code":    strconv.Itoa(http.StatusBadRequest),
			"message": err.Error(),
		}

		c.Context.AbortWithStatusJSON(http.StatusBadRequest, h)

		return
	}

	c.Context.JSON(http.StatusNoContent, gin.H{})
}
