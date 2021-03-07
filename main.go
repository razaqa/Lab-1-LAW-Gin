package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.POST("/calculate", calculateAB)
	router.Run(":8080")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func calculateAB(c *gin.Context) {

	var payload struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := payload.A + payload.B

	c.JSON(http.StatusOK, gin.H{
		"a":     payload.A,
		"b":     payload.B,
		"hasil": result,
	})
}
