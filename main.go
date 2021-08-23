package main

import (
	"net/http"

	"github.com/bwhour/sep31-demo/header"
	"github.com/gin-gonic/gin"
)

type info struct {
	Amount      uint64 `json:"amount"`
	Asset_code  string `json:"asset_code"`
	Sender_id   string `json:"sender_id"`
	Receiver_id string `json:"receiver_id"`
	Field       map[string]string
}

type transaction struct {
	Id                 string `json:"id"`
	Stellar_account_id string `json:"stellar_account_id"`
	Stellar_memo_type  string `json:"stellar_memo_type"`
	Stellar_memo       string `json:"stellar_memo"`
	Extra_info         map[string]string
	How                string `json:"how"`
}

type transinfo struct {
	Id                      string `json:"id"`
	Status                  string `json:"status"`
	Amount_in               string `json:"amount_in"`
	Amount_out              string `json:"amount_out"`
	Amount_fee              string `json:"amount_fee"`
	Stellar_account_id      string `json:"stellar_account_id"`
	Stellar_memo_type       string `json:"stellar_memo_type"`
	Stellar_memo            string `json:"stellar_memo"`
	Started_at              int64  `json:"timestamp"`
	Stellar_transaction_id  string `json:"stellar_transaction_id"`
	External_transaction_id string `json:"external_transaction_id"`
}

func main() {
	r := gin.Default()

	r.Use(header.NoCache)
	r.Use(header.Options)
	r.Use(header.Secure)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
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
