package itemmodels

type LookForItemInput struct {
	//SellerUserID    int    `json:"User ID"`
	//ItemID    int    `json:"Item ID"`

	ItemCategory      string `json:"Item Category"`
	ItemImageLinkURL  string `json:"Item ImageURL,omitempty"`
	ItemName          string `json:"Item Name"`
	ItemPreferedPrice int    `json:"Item Preferred Price"`
	ItemQty           int    `json:"Item Quantity"`
	ItemDesc          string `json:"Item Description"`
}
