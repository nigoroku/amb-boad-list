package controller

import (
	"fmt"

	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FindInputTimeline(c *gin.Context) {

	timelineService := service.NewTimelineService()
	timeline, err := timelineService.FindInputTimeline()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"timeline": timeline,
	})
}

func FindOutputTimeline(c *gin.Context) {

	timelineService := service.NewTimelineService()
	timeline, err := timelineService.FindOutputTimeline()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"timeline": timeline,
	})
}
