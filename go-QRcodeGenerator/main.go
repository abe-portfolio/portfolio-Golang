package main

import (
	"log"

	"go-QRcode/M_Generator"
	"go-QRcode/M_Reader"

	"github.com/gin-gonic/gin"
)

func main() {
	// 他モジュールの呼び出し確認
	M_Reader.R_index()
	M_Generator.G_index()

	r := gin.Default()

	// HTMLテンプレートをロード
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	// ロードされたテンプレートをログ出力して確認
	log.Println("Loaded HTML Templates:")
	r.HTMLRender.Instance("dummy", nil)

	// ルートエンドポイント
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "top_page.html", gin.H{"name": "Gin-User"})
	})

	// サーバー起動
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
