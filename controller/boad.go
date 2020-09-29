package controller

import (
	"fmt"

	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBoadList(c *gin.Context) {

	boadService := service.NewBoadService()
	users, err := boadService.FindBoadList()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "ok",
		"boad_list": users,
	})
}
