package types

type Product struct {
	StockCode    string  `json:"stockCode"`
	Description  string  `json:"description"`
	ItemType     int     `json:"itemType"`
	NominalCode  string  `json:"nominalCode"`
	UnitOfSale   string  `json:"unitOfSale"`
	DeletedFlag  bool    `json:"deletedFlag"`
	InactiveFlag bool    `json:"inactiveFlag"`
	SalesPrice   float64 `json:"salesPrice"`
	QtyAllocated float64 `json:"qtyAllocated"`
	QtyInStock   float64 `json:"qtyInStock"`
}
type ProductResponse struct {
	Success  bool `json:"success"`
	Code     int  `json:"code"`
	Response Product
	Message  string `json:"message"`
}
