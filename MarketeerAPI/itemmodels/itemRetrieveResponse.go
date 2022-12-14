package itemmodels

type ItemRetrieveResponse struct {
	Code         int32  `json:"code,omitempty"`
	Status       string `json:"status,omitempty"`
	Message      string `json:"message,omitempty"`
	ItemID       int    `json:"Item ID,omitempty"`
	ItemCategory string `json:"Item Caegory"`
	SellerID     int    `json:"Seller ID"`

	// FirstName  string `json:"FirstName"`
	// LastName   string `json:"LastName"`
	// EMail      string `json:"E-Mail"`
	// ContactNum string `json:"ContactNumber"`
	// BirthDate  string `json:"BirthDate"`
	// Address    string `json:"Address"`
	// UserName   string `json:"Username"`
	
	ItemLookID       int    `json:"Item Look ID,omitempty"`
	ItemImageLinkURL string `json:"Item Image URL,omitempty"`
	ItemName         string `json:"Item Name,omitempty"`
	ItemPrefPrice    int    `json:"Item Preferred Price,omitempty"`
	ItemPrice        int    `json:"Item Price,omitempty"`
	ItemQty          int    `json:"Item Quantity,omitempty"`
	ItemDesc         string `json:"Item Description,omitempty"`


	
}
