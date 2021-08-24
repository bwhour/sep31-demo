package main

import (
	"net/http"

	"github.com/bwhour/sep31-demo/header"
	"github.com/bwhour/sep31-demo/schema"
	"github.com/gin-gonic/gin"
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
	r.GET("/sep0031/transactions/:id", getTransaction)
	r.POST("/sep0031/transactions", createTransaction)

	r.Run(":9935") // listen and serve on 0.0.0.0:9935
}

func getInfo(c *gin.Context) {
    resp := map[string]interface{} {
		    "receive":map[string]interface{} {
			    "USDC":map[string]interface{} {
				    "quotes_supported": true,
				    "fee_fixed": 5,
				    "fee_percent": 1,
				    "min_amount": 0.1,
				    "max_amount": 1000,
				    "sep12": map[string]interface{}{
					    "sender":map[string]interface{} {
						    "types":map[string]interface{} {
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
				    "fields":map[string]interface{}{
					    "transaction":map[string]interface{}{
						    "receiver_routing_number":map[string]interface{}{
							    "description": "routing number of the destination bank account",
						    },
						    "receiver_account_number":map[string]interface{}{
							    "description": "bank account number of the destination",
						    },
						    "type":map[string]interface{}{
							    "description": "type of deposit to make",
							    "choices":[]string{"SEPA","SWIFT"},
						    },
					    },
				    },
			    },
	    },
	
	
    }
	c.JSON(http.StatusOK, gin.H{
		"method": "getInfo",
		"resp":resp,
	})
}

func getTransaction(c *gin.Context) {
	id := c.Param("id")
	resp := []schema.Transinfo{
		{
			Id:                      "3bf73ba8-8a48-41b7-99da-b3d70ad959c0",
			Status:                  "completed",
			Amount_in:               "53.0000000",
			Amount_out:              "51.6066212",
			Amount_fee:              "0.0000000",
			Stellar_account_id:      "GBUQO65XW7TDXIHND7MFNIUZU6UNBA5XN27ODFDXVOLZYOXQSQUTPVHW",
			Stellar_memo_type:       "text",
			Stellar_memo:            "psp:2850",
			Started_at:              "2020-10-06 14:50:40.508155 +0000 UTC",
			Stellar_transaction_id:  "18b6a4b464b061aefd949c00645deb6732b18ed3acb723d3ba5ac6ba00465d94",
			External_transaction_id: "18b6a4b464b061aefd949c00645deb6732b18ed3acb723d3ba5ac6ba00465d94",
		}}
	c.JSON(http.StatusOK, gin.H{
		"method": "getTransactions",
		"id":     id,
		"resp":   resp,
	})
}

func createTransaction(c *gin.Context) {
	resp := []schema.Transaction{
		{
			Id:                 "6238dd28-a47a-470e-9cee-7f547730d6b1",
			Stellar_account_id: "GBUQO65XW7TDXIHND7MFNIUZU6UNBA5XN27ODFDXVOLZYOXQSQUTPVHW",
			Stellar_memo_type:  "text",
			Stellar_memo:       "psp:2852",
			Extra_info: map[string]interface{}{
				"amount":                  "53",
				"bank_name":               "BRED BAGNOLET",
				"beneficiary":             "TEMPO FRANCE",
				"beneficiary_address":     "89 BOULEVARD DE MAGENTA 75010 PARIS",
				"bic":                     "BREDFRPPXXX",
				"currency":                "EUR",
				"external_transaction_id": "5b22d1ed-b789-48e0-9d45-9d2f116f0159",
				"iban":                    "FR76 1010 7002 3000 5220 3386 959",
				"reference_id":            "swift_101",
				"transaction_id":          "5b22d1ed-b789-48e0-9d45-9d2f116f0159",
			},
			How: "SWIFT",
		}}
	c.JSON(http.StatusOK, gin.H{
		"method": "createTransaction",
		"resp":   resp,
	})
}
