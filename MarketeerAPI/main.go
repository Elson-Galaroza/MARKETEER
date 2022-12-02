package main

import (
	"marketeer/api"
	"marketeer/displaydataapi"
	"marketeer/itemsapi"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// System
	router.GET("/marketeer/systemcheck", api.SystemCheck)
	// User Registration and Login
	router.POST("/marketeer/registerUser", api.RegisterUser)
	router.POST("/marketeer/login", api.LogInUser)
	//sell item
	router.POST("/marketeer/sellitem", itemsapi.InputSellItemDetails)
	//display all item for sale
	router.GET("/marketeer/displayitemforsale", displaydataapi.RetrieveForSaleItems)
	//sell item
	router.POST("/marketeer/lookforitem", itemsapi.InputItemsForLookDetails)
	//display all item for look
	router.POST("/marketeer/searchItem", itemsapi.ItemSearch)
	router.Run(":80")
}
