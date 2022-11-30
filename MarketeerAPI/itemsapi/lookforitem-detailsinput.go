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

func InputItemsForLookDetails(c *gin.Context) {

	var LookItemPayload itemmodels.LookForItemInput
	var genereicResponse models.GenericResponse
	var itemRetrieveResponse itemmodels.ItemRetrieveResponse

	errJson := c.ShouldBindJSON(&LookItemPayload)

	if errJson == nil {

		itemRetrieveResponse.Code = 200
		itemRetrieveResponse.Status = "Succesfull Retrieve"

		itemRetrieveResponse.SellerID = 122 //chill sa
		itemRetrieveResponse.ItemImageLinkURL = LookItemPayload.ItemImageLinkURL
		itemRetrieveResponse.ItemCategory = LookItemPayload.ItemCategory
		itemRetrieveResponse.ItemName = LookItemPayload.ItemName
		itemRetrieveResponse.ItemPrefPrice = LookItemPayload.ItemPreferedPrice
		itemRetrieveResponse.ItemQty = LookItemPayload.ItemQty
		itemRetrieveResponse.ItemDesc = LookItemPayload.ItemDesc

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

	LookerID := 1
	ItemCategory := LookItemPayload.ItemCategory
	ItemImageLinkURL := LookItemPayload.ItemImageLinkURL
	ItemName := LookItemPayload.ItemName
	ItemPrefPrice := LookItemPayload.ItemPreferedPrice
	ItemQty := LookItemPayload.ItemQty
	ItemDesc := LookItemPayload.ItemDesc

	insertStatement := " INSERT INTO `marketeer_database1`.`ItemsForLooked` ( `ItemCategory` ,`LookerID`, `ItemImageLinkURL`, `ItemName`, `ItemPreferredPrice`, `ItemQty`, `ItemDesc`) VALUES (?, ?, ?, ?, ?, ?, ?); "
	insert, err := db.Query(insertStatement, ItemCategory, LookerID, ItemImageLinkURL, ItemName, ItemPrefPrice, ItemQty, ItemDesc)

	if err != nil {
		panic(err.Error())

	}
	defer insert.Close()
	fmt.Print("Successfull Connection")

	// var (
	// 	itemLookID    int
	// 	itemCategory  string
	// 	lookerID      int
	// 	imageURL      string
	// 	itemName      string
	// 	itemPrefPrice int
	// 	itemQty       int
	// 	itemDesc      string
	// )
	// getItemIDStatement := "SELECT * FROM marketeer_database1.ItemsForLooked ORDER BY LookID DESC LIMIT 1"

	// row := db.QueryRow(getItemIDStatement)
	// err2 := row.Scan()

	// if err2 != nil && err2 != sql.ErrNoRows {
	// 	rows, err := db.Query(getItemIDStatement)
	// 	for rows.Next() {
	// 		err := rows.Scan(&itemLookID, &itemCategory, &lookerID, &imageURL, &itemName, &itemPrefPrice, &itemQty, &itemDesc)
	// 		if err != nil && err != sql.ErrNoRows {
	// 			print("\nError\n")
	// 			log.Fatal(err)
	// 		}
	// 	}
	// 	if err == nil && err != sql.ErrNoRows {
	// 		println("\n", itemLookID)
	// 	}
	// }

	// //itemIDnum := itemID

	// itemRetrieveResponse.ItemLookID = itemLookID

}
