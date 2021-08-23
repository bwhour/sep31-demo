package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// add router.
	r.GET("/info/", getInfo)
	r.GET("/transaction", getTransaction)
	r.POST("/transactions", createTransaction)

	r.Run(":9935") // listen and serve on 0.0.0.0:9935
}

func getInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "getInfo",
	})
}

func getTransaction(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": "getTransaction",
	})
}

func createTransaction(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": "createTransaction",
	})
}
