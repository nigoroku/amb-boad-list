package main

import (
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
			v1.GET("/input/selected_actions", controller.FindInputSelectedActions)
			v1.GET("/output/selected_actions", controller.FindOutputSelectedActions)
			v1.POST("/input/update_action", controller.UpdateInputAction)
			v1.POST("/output/update_action", controller.UpdateOutputAction)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run(":8084")
}
