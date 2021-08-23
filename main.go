package main

import (
	"net/http"

	"github.com/bwhour/sep31-demo/header"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(header.NoCache)
	r.Use(header.Options)
	r.Use(header.Secure)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// add router.
	r.GET("/sep0031/info/", getInfo)
	r.GET("/sep0031/transactions/:id", getTransaction)
	r.POST("/sep0031/transactions", createTransaction)

	r.Run(":9935") // listen and serve on 0.0.0.0:9935
}

func getInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "getInfo",
	})
}

func getTransaction(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"data": "getTransactions",
		"id":   id,
	})
}

func createTransaction(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": "createTransaction",
	})
}
