package main

import (
	"net/http"

	"github.com/annts095/gin-practice/controller"
	"github.com/annts095/gin-practice/database"
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
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("items", func(c *gin.Context) {
		itemController := controller.ItemController{Context: c}
		itemController.Create()
	})
	r.GET("items_migrate", func(c *gin.Context) {
		database.ItemMigrate()
		c.JSON(http.StatusOK, gin.H{
			"message": "migratie success: items",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
