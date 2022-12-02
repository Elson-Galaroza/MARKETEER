package itemsapi

import (
	"database/sql"
	"fmt"

	"log"
	"marketeer/config"
	"marketeer/itemmodels"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func ItemSearch(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var searchItemPayload itemmodels.SearchItems
	var retrieveResponse itemmodels.ItemRetrieveResponse

	err := c.ShouldBindJSON(&searchItemPayload)

	cfg := mysql.Config{
		User:   config.GetEnvConfig("DBUSER"),
		Passwd: config.GetEnvConfig("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "marketeer_database1",
	}

	var dberr error
	db, dberr = sql.Open("mysql", cfg.FormatDSN())
	if dberr != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err)
	}

	searchItemForSaleStatement := "SELECT `ItemID`, `ItemCategory`, `SellerID`, `ItemImageLinkURL`, `ItemName`, `ItemPrice`, `ItemQty`, `ItemDesc` FROM marketeer_database1.ItemsForSale WHERE `ItemName` = '?';"
	//searchItemForSaleStatement := "SELECT FROM `marketeer_database1`.`ItemsForSale` (`ItemID`, `ItemCategory`, `SellerID`, `ItemImageLinkURL`, `ItemName`, `ItemPrice`, `ItemQty`, `ItemDesc`) WHERE `ItemName`=  ?;"
	row := db.QueryRow(searchItemForSaleStatement, searchItemPayload.ItemNameSearch)
	err2 := row.Scan()

	var (
		itemID       int
		itemCategory string
		sellerID     int
		itemImageURL string
		itemName     string
		itemPrice    int
		itemQty      int
		itemDesc     string
	)
	if err2 != nil && err2 != sql.ErrNoRows {
		print("\nTHERE IS THAT USER \n")

		rows, err := db.Query(searchItemForSaleStatement)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&itemID, &itemCategory, &sellerID, &itemImageURL, &itemName, &itemPrice, &itemQty, &itemDesc)
			if err != nil {
				log.Fatal(err)

			}

		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	retrieveResponse.ItemCategory = itemCategory
	retrieveResponse.SellerID = sellerID
	retrieveResponse.ItemImageLinkURL = itemImageURL
	retrieveResponse.ItemID = itemID
	retrieveResponse.ItemName = itemName
	retrieveResponse.ItemPrice = itemPrice
	retrieveResponse.ItemQty = itemQty
	retrieveResponse.ItemDesc = itemDesc
	//log.Println(itemID, itemName)
	c.JSON(http.StatusOK, retrieveResponse)

	currentTime := time.Now()
	fmt.Println("The time is", currentTime)
}
