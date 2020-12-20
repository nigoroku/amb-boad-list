package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"local.packages/models"
	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ActionJson struct {
	UserID        int    `json:"user_id"`
	ActionType    string `json:"action_type"`
	Insert        bool   `json:"insert"`
	AchievementID int    `json:"achievement_id"`
}

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

	j, _ := json.Marshal(&timeline)
	var decoded []models.Timeline
	json.Unmarshal([]byte(string(j)), &decoded)

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"timeline": decoded,
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

	j, _ := json.Marshal(&timeline)
	var decoded []models.Timeline
	json.Unmarshal([]byte(string(j)), &decoded)

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"timeline": timeline,
	})
}

// FindInputSelectedActions インプット実績に対する選択済みのアクションを取得する
func FindInputSelectedActions(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	timelineService := service.NewTimelineService()
	actions, err := timelineService.FindInputSelectedActions(userID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "ok",
		"selected_actions": actions,
	})
}

// FindOutputSelectedActions アウトプット実績に対する選択済みのアクションを取得する
func FindOutputSelectedActions(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	timelineService := service.NewTimelineService()
	actions, err := timelineService.FindOutputSelectedActions(userID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "ok",
		"selected_actions": actions,
	})
}

// UpdateInputAction ユーザーに紐づくインプット実績を更新する
func UpdateInputAction(c *gin.Context) {

	var actionJSON ActionJson
	if err := c.ShouldBindJSON(&actionJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timelineService := service.NewTimelineService()
	err := timelineService.UpdateInputAction(actionJSON.UserID, actionJSON.ActionType, actionJSON.Insert, actionJSON.AchievementID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// UpdateOutputAction ユーザーに紐づくインプット実績を更新する
func UpdateOutputAction(c *gin.Context) {

	var actionJSON ActionJson
	if err := c.ShouldBindJSON(&actionJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timelineService := service.NewTimelineService()
	err := timelineService.UpdateOutputAction(actionJSON.UserID, actionJSON.ActionType, actionJSON.Insert, actionJSON.AchievementID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
