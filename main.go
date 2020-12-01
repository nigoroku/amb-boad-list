package main

import (
	"os"

	"github.com/gin-contrib/cors"

	// "github.com/kzpolicy/user/controller"
	// "github.com/kzpolicy/user/middleware"
	"local.packages/controller"
	"local.packages/db"
	"local.packages/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --output models/generated --pkgname generated --wipe mysql

func main() {
	r := gin.Default()

	// ミドルウェア
	r.Use(middleware.RecordUaAndTime)

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// DB接続
	db.Init()

	// ルーティング
	TimelineRoute := r.Group("/api/v1")
	{
		v1 := TimelineRoute.Group("/timeline")
		{
			v1.GET("/input", controller.FindInputTimeline)
			v1.GET("/output", controller.FindOutputTimeline)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// 起動ポートを環境変数から取得
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8084"
	}

	r.Run(":" + port)
}
