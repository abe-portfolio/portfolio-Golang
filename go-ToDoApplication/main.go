package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/labstack/echo/v4"  Echo
	// GORMドライバー：SQLite
)

func main() {
	// 動作確認用
	fmt.Println("Hello, World")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"name": "Gin-User",
		})
	})

	r.Run(":8080")
}
