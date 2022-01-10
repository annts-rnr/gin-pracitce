package main

import (
	"net/http"
	"strconv"

	"github.com/annts095/gin-practice/database"
	"github.com/annts095/gin-practice/input"
	"github.com/annts095/gin-practice/model"
	"github.com/annts095/gin-practice/repository"
	"github.com/gin-gonic/gin"
)

type H map[string]interface{}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "assets/css")
	r.Static("/js", "assets/js")
	r.GET("index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main websites",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("items", func(c *gin.Context) {
		var input input.ItemCreateInput

		if err := c.ShouldBindJSON(&input); err != nil {
			h := map[string]string{
				"code":    strconv.Itoa(http.StatusBadRequest),
				"message": err.Error(),
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, h)

			return
		}
		item := model.Item{
			Title:    input.Item.Title,
			Contents: input.Item.Contents,
			Price:    input.Item.Price,
		}

		repository := repository.ItemRepository{DB: database.GetGormConnect()}
		repository.Save(&item)
		c.JSON(200, gin.H{
			"id":       item.ID,
			"title":    item.Title,
			"contents": item.Contents,
			"price":    item.Price,
		})
	})
	r.GET("items_migrate", func(c *gin.Context) {
		database.ItemMigrate()
		c.JSON(200, gin.H{
			"message": "migratie success: items",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
