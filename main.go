package main

import (
	"net/http"
	"strings"
	
	"github.com/bwhour/sep31-demo/header"
	"github.com/bwhour/sep31-demo/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
	r.GET("/sep31/info/", getInfo)
	r.GET("/sep12/customer/", getCustomer)
	r.GET("/sep31/transactions/:id", getTransaction)
	r.PUT("/sep12/customer", updateCustomer)
	r.POST("/sep31/transactions", createTransaction)
	
	r.Run(":9935") // listen and serve on 0.0.0.0:9935
}

func getInfo(c *gin.Context) {
	resp := map[string]interface{}{
		"receive": map[string]interface{}{
			"VUSD": map[string]interface{}{
				"enabled": true,
				"sep_12": map[string]interface{}{
					"sender": map[string]interface{}{
						"type": map[string]interface{}{
							"sep31-sender": map[string]interface{}{
								"description": "sender",
							},
						},
					},
					"receiver": map[string]interface{}{
						"type": map[string]interface{}{
							"sep31-receiver": map[string]interface{}{
								"description": "receiver",
							},
						},
					},
				},
				"fields": map[string]interface{}{
					"transaction": map[string]interface{}{
						"receiver_account_number": map[string]interface{}{
							"type":        "string",
							"description": "Receiver account number where payout will be executed to."},
					},
				},
			},
		},
	}
	c.JSON(http.StatusOK, resp)
}

func getTransaction(c *gin.Context) {
	id := c.Param("id")
	resp := schema.Transinfo{
		Id:               id,
		Status:           "completed",
		StellarAccountId: "GDRZFOASKJUV5E2ETWJY5RZRKQF3PMCZ2CNLLWH365X7OUM5NHUOVE7I",
		StellarMemoType:  "text",
		StellarMemo:      "velo:GMWJGSAPGATL",
	}
	c.JSON(http.StatusOK, resp)
}

func getCustomer(c *gin.Context) {
	
	url_type := c.Query("type")
	url_id := c.Query("id")
	url_memo := c.Query("memo")
	if strings.EqualFold(url_type, "sep31-sender") ||
		strings.EqualFold(url_id, "be68fbb2-0aa1-408b-9728-e2259a352a82") ||
		strings.EqualFold(url_memo, "email:ya.gekpp@yandex.ru") {
		url_id = "be68fbb2-0aa1-408b-9728-e2259a352a82"
	} else if strings.EqualFold(url_type, "sep31-receiver") ||
		strings.EqualFold(url_id, "9114353d-4b51-47dc-ae19-372b9d43fffa") ||
		strings.EqualFold(url_memo, "email:lada@bitazza.com") {
		url_id = "9114353d-4b51-47dc-ae19-372b9d43fffa"
	}
	resp := schema.CustomerInfo{
		// The case when a customer has been successfully KYC'd and approved
		Status: "NEEDS_INFO",
		Fields: map[string]interface{}{
			"email": map[string]interface{}{
				"description": "The customer's email address",
				"type":        "string",
			},
			"first_name": map[string]interface{}{
				"description": "The customer's first name",
				"type":        "string",
			},
			"last_name": map[string]interface{}{
				"description": "The customer's last name",
				"type":        "string",
			},
			"mobile_number": map[string]interface{}{
				"description": "The customer's phone number",
				"type":        "string",
			},
		},
	}
	
	c.JSON(http.StatusOK, resp)
}

func updateCustomer(c *gin.Context) {
	var id string
	url_type := c.PostForm("type")
	
	if strings.EqualFold(url_type, "sep31-sender") {
		id = "be68fbb2-0aa1-408b-9728-e2259a352a82"
	} else if strings.EqualFold(url_type, "sep31-receiver") {
		id = "9114353d-4b51-47dc-ae19-372b9d43fffa"
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "status": "ACCEPTED"})
}

func createTransaction(c *gin.Context) {
	id := uuid.New().String()
	resp := schema.Transaction{
		Id:                 id,
		Stellar_account_id: "GDRZFOASKJUV5E2ETWJY5RZRKQF3PMCZ2CNLLWH365X7OUM5NHUOVE7I",
		Stellar_memo_type:  "text",
		Stellar_memo:       "velo:WWMNPPQIMGEG",
	}
	c.JSON(http.StatusOK, resp)
}