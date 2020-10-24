package types

import "time"

type Adjustments struct {
	Id          int       `db:"id"`
	Adjustment  string    `db:"adjustment"`
	Amount      float64   `db:"amount"`
	StockCode   string    `db:"stock_code"`
	Reference   string    `db:"reference_text"`
	SageUpdated int       `db:"sage_updasted"`
	InsertedAt  time.Time `db:"inserted_at"`
}
