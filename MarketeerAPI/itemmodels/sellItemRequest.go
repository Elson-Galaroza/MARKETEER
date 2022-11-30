package itemmodels

type SellItemInput struct {
	//SellerUserID    int    `json:"User ID"`
	//ItemID    int    `json:"Item ID"`
	ItemCategory     string `json:"Item Category"`
	ItemImageLinkURL string `json:"Item ImageURL"`
	ItemName         string `json:"Item Name"`
	ItemPrice        int    `json:"Item Price"`
	ItemQty          int    `json:"Item Quantity"`
	ItemDesc         string `json:"Item Description"`
}
