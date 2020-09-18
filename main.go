package main

import (
	"net/http"

	"bubblebot.service/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run(":8093")
}
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/getUser", func(c *gin.Context) {
		vals := &model.OrderResponse{
			UserID:   "20000",
			Username: "112312",
			Password: "WakeUp",
			Broker:   "008",
		}

		if err := c.ShouldBindJSON(&vals); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "InternalServerError"})
			return
		}
		c.JSON(http.StatusOK, vals)

	})

	return router
}
