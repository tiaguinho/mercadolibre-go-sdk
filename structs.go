package meli

//structs for test methods only
//to use the api you have to implement your on structs

//Category struct
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Product struct
type Product struct {
	ID                string  `json:"id,omitempty"`
	ListingTypeID     string  `json:"listing_type_id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	CategoryID        string  `json:"category_id"`
	BuyingMode        string  `json:"buying_mode"`
	CurrencyID        string  `json:"currency_id"`
	Condition         string  `json:"condition"`
	Price             float32 `json:"price"`
	AvailableQuantity int32   `json:"available_quantity"`
	Pictures          []Image `json:"pictures"`
}

//Image struct
type Image struct {
	Source string `json:"source"`
}

//Status struct
type Status struct {
	Status string `json:"status"`
}
