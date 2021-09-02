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
	r.GET("/sep0031/info/", getInfo)
	r.GET("/sep0012/customer/", getCustomer)
	r.GET("/sep0031/transactions/:id", getTransaction)
	r.PUT("/sep0012/customer", updateCustomer)
	r.POST("/sep0031/transactions", createTransaction)
	
	r.Run(":9935") // listen and serve on 0.0.0.0:9935
}

func getInfo(c *gin.Context) {
	resp := map[string]interface{}{
		"receive": map[string]interface{}{
			"VUSD": map[string]interface{}{
				"quotes_supported": true,
				"fee_fixed":        5,
				"fee_percent":      1,
				"min_amount":       0.1,
				"max_amount":       1000,
				"sep12": map[string]interface{}{
					"sender": map[string]interface{}{
						"types": map[string]interface{}{
							"sep31-sender": map[string]interface{}{
								"description": "U.S. citizens limited to sending payments of less than $10,000 in value ",
							},
							"sep31-large-sender": map[string]interface{}{
								"description": "U.S. citizens that do not have sending limits",
							},
							"sep31-foreign-sender": map[string]interface{}{
								"description": "non-U.S. citizens sending payments of less than $10,000 in value",
							},
						},
					},
					"receiver": map[string]interface{}{
						"types": map[string]interface{}{
							"sep31-receiver": map[string]interface{}{
								"description": "U.S. citizens receiving USD",
							},
							"sep31-foreign-receiver": map[string]interface{}{
								"description": "non-U.S. citizens receiving USD",
							},
						},
					},
				},
				"fields": map[string]interface{}{
					"transaction": map[string]interface{}{
						"receiver_routing_number": map[string]interface{}{
							"description": "routing number of the destination bank account",
						},
						"receiver_account_number": map[string]interface{}{
							"description": "bank account number of the destination",
						},
						"type": map[string]interface{}{
							"description": "type of deposit to make",
							"choices":     []string{"SEPA", "SWIFT"},
						},
					},
				},
			},
		},
	}
	c.JSON(http.StatusOK, resp)
}

func getTransaction(c *gin.Context) {
	id := c.Param("id")
	resp := []schema.Transinfo{
		{
			Id:                    id,
			Status:                "completed",
			AmountIn:              "53.0000000",
			AmountOut:             "51.6066212",
			AmountFee:             "0.0000000",
			StellarAccountId:      "GDRZFOASKJUV5E2ETWJY5RZRKQF3PMCZ2CNLLWH365X7OUM5NHUOVE7I",
			StellarMemoType:       "text",
			StellarMemo:           "psp:2850",
			StartedAt:             "2020-10-06 14:50:40.508155 +0000 UTC",
			StellarTransactionId:  "18b6a4b464b061aefd949c00645deb6732b18ed3acb723d3ba5ac6ba00465d94",
			ExternalTransactionId: "18b6a4b464b061aefd949c00645deb6732b18ed3acb723d3ba5ac6ba00465d94",
		},
	}
	c.JSON(http.StatusOK, resp)
}

func getCustomer(c *gin.Context) {
	
	url_type := c.Query("type")
	url_id := c.Query("id")
	url_memo := c.Query("memo")
	if strings.EqualFold(url_type, "sender") ||
		strings.EqualFold(url_id, "391fb415-c223-4608-b2f5-dd1e91e3a986") ||
		strings.EqualFold(url_memo, "email:ya.gekpp@yandex.ru") {
		url_id = "391fb415-c223-4608-b2f5-dd1e91e3a986"
	} else if strings.EqualFold(url_type, "receiver") ||
		strings.EqualFold(url_id, "391fb415-c223-4608-b2f5-dd1e91e3a999") ||
		strings.EqualFold(url_memo, "email:lada@bitazza.com") {
		url_id = "391fb415-c223-4608-b2f5-dd1e91e3a999"
	}
	resp := schema.CustomerInfo{
		// The case when a customer has been successfully KYC'd and approved
		Id:     url_id,
		Status: "ACCEPTED",
		Provided_fields: map[string]interface{}{
			"first_name": map[string]interface{}{
				"description": "The customer's first name",
				"type":        "string",
				"status":      "ACCEPTED",
			},
			"last_name": map[string]interface{}{
				"description": "The customer's last name",
				"type":        "string",
				"status":      "ACCEPTED",
			},
			"email_address": map[string]interface{}{
				"description": "The customer's email address",
				"type":        "string",
				"status":      "ACCEPTED",
			},
			"phone_number": map[string]interface{}{
				"description": "The customer's phone number",
				"status":      "ACCEPTED",
				"type":        "string",
			},
		},
	}
	
	c.JSON(http.StatusOK, resp)
}

func updateCustomer(c *gin.Context) {
	var id string
	url_type := c.PostForm("type")
	
	if strings.EqualFold(url_type, "SENDER") {
		id = "391fb415-c223-4608-b2f5-dd1e91e3a986"
	} else if strings.EqualFold(url_type, "RECEIVER") {
		id = "391fb415-c223-4608-b2f5-dd1e91e3a999"
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func createTransaction(c *gin.Context) {
	id := uuid.New().String()
	resp := []schema.Transaction{
		{
			Id:                 id,
			Stellar_account_id: "GDRZFOASKJUV5E2ETWJY5RZRKQF3PMCZ2CNLLWH365X7OUM5NHUOVE7I",
			Stellar_memo_type:  "text",
			Stellar_memo:       "psp:18000",
			Extra_info: map[string]interface{}{
				"receiver_asset_code": "USDC",
				"receiver_asset_amt":  "23.43",
				"exchange_pair":       "USDC/THB",
				"exchange_rate":       "32.931",
				"buy_or_sell":         "sell",
			},
		},
	}
	c.JSON(http.StatusOK, resp)
}