package itemsapi

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"log"
	"marketeer/config"
	"marketeer/itemmodels"
	"marketeer/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
)

var db *sql.DB

func InputSellItemDetails(c *gin.Context) {

	var sellItemPayload itemmodels.SellItemInput
	var genereicResponse models.GenericResponse
	var itemRetrieveResponse itemmodels.ItemRetrieveResponse

	errJson := c.ShouldBindJSON(&sellItemPayload)

	if errJson == nil {

		itemRetrieveResponse.Code = 200
		itemRetrieveResponse.Status = "Succesfull Retrieve"

		itemRetrieveResponse.SellerID = 122 //chill sa
		itemRetrieveResponse.ItemImageLinkURL = sellItemPayload.ItemImageLinkURL
		itemRetrieveResponse.ItemCategory = sellItemPayload.ItemCategory
		itemRetrieveResponse.ItemName = sellItemPayload.ItemName
		itemRetrieveResponse.ItemPrice = sellItemPayload.ItemPrice
		itemRetrieveResponse.ItemQty = sellItemPayload.ItemQty
		itemRetrieveResponse.ItemDesc = sellItemPayload.ItemDesc

		// retrieveResponse.FirstName = registrationPayload.FirstName
		// retrieveResponse.LastName = registrationPayload.LastName
		// retrieveResponse.EMail = registrationPayload.EMail

		// retrieveResponse.ContactNum = registrationPayload.ContactNum
		// retrieveResponse.BirthDate = registrationPayload.BirthDate
		// retrieveResponse.Address = registrationPayload.Address
		// retrieveResponse.UserName = registrationPayload.UserName
		// retrieveResponse.Password = registrationPayload.Password
	} else {
		genereicResponse.Status = "Failed"
		genereicResponse.Code = 500
		genereicResponse.Message = errJson.Error()

	}

	c.JSON(http.StatusOK, itemRetrieveResponse)

	cfg := mysql.Config{
		User:   config.GetEnvConfig("DBUSER"),
		Passwd: config.GetEnvConfig("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "marketeer_database1",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	SellerID := 1
	ItemCategory := sellItemPayload.ItemCategory
	ItemImageLinkURL := sellItemPayload.ItemImageLinkURL
	ItemName := sellItemPayload.ItemName
	ItemPrice := sellItemPayload.ItemPrice
	ItemQty := sellItemPayload.ItemQty
	ItemDesc := sellItemPayload.ItemDesc

	insertStatement := " INSERT INTO `marketeer_database1`.`ItemsForSale` ( `ItemCategory` ,`SellerID`, `ItemImageLinkURL`, `ItemName`, `ItemPrice`, `ItemQty`, `ItemDesc`) VALUES (?, ?, ?, ?, ?, ?, ?); "
	insert, err := db.Query(insertStatement, ItemCategory, SellerID, ItemImageLinkURL, ItemName, ItemPrice, ItemQty, ItemDesc)

	if err != nil {
		panic(err.Error())

	}
	defer insert.Close()
	fmt.Print("Successfull Connection")

	var (
		itemID       int
		itemCategory string
		sellerID     int
		imageURL     string
		itemName     string
		itemPrice    int
		itemQty      int
		itemDesc     string
	)
	getItemIDStatement := "SELECT * FROM marketeer_database1.ItemsForSale ORDER BY ItemID DESC LIMIT 1"

	row := db.QueryRow(getItemIDStatement)
	err2 := row.Scan()

	if err2 != nil && err2 != sql.ErrNoRows {
		rows, err := db.Query(getItemIDStatement)
		for rows.Next() {
			err := rows.Scan(&itemID, &itemCategory, &sellerID, &imageURL, &itemName, &itemPrice, &itemQty, &itemDesc)
			if err != nil && err != sql.ErrNoRows {
				print("\nError\n")
				log.Fatal(err)
			}
		}
		if err == nil && err != sql.ErrNoRows {
			println("\n", itemID)
		}
	}

	//itemIDnum := itemID

	itemRetrieveResponse.ItemID = itemID

}
