package types

type Adjustments struct {
	Id             int     `db:"id"`
	AdjustmentType string  `db:"adjustment_type"`
	Amount         float64 `db:"amount"`
	StockCode      string  `db:"stock_code"`
	Reference      string  `db:"reference_text"`
}
