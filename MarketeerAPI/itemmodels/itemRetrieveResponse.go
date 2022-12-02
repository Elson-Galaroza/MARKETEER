package itemmodels

type ItemRetrieveResponse struct {
	Code         int32  `json:"code,omitempty"`
	Status       string `json:"status,omitempty"`
	Message      string `json:"message,omitempty"`
	ItemID       int    `json:"ItemID,omitempty"`
	ItemCategory string `json:"ItemCategory"`
	SellerID     int    `json:"SellerID"`

	// FirstName  string `json:"FirstName"`
	// LastName   string `json:"LastName"`
	// EMail      string `json:"E-Mail"`
	// ContactNum string `json:"ContactNumber"`
	// BirthDate  string `json:"BirthDate"`
	// Address    string `json:"Address"`
	// UserName   string `json:"Username"`

	ItemLookID       int    `json:"ItemLookID,omitempty"`
	ItemImageLinkURL string `json:"ItemImageURL,omitempty"`
	ItemName         string `json:"ItemName,omitempty"`
	ItemPrefPrice    int    `json:"ItemPreferred Price,omitempty"`
	ItemPrice        int    `json:"ItemPrice,omitempty"`
	ItemQty          int    `json:"ItemQuantity,omitempty"`
	ItemDesc         string `json:"ItemDescription,omitempty"`
}
