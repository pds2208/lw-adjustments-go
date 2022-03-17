package types

import (
	"time"
)

type Adjustments struct {
	Id             int        `db:"id,omitempty"`
	AdjustmentType string     `db:"adjustment_type"`
	Amount         float64    `db:"amount"`
	StockCode      string     `db:"stock_code"`
	AdjustmentDate *time.Time `db:"adjustment_date"`
	Batch          string     `db:"batch"`
	SageUpdated    *bool      `db:"sage_updated"`
	InsertedAt     *time.Time `db:"inserted_at"`
	SageUpdatedAt  *time.Time `db:"sage_updated_at"`
	NumRetries     int        `db:"num_retries"`
	UpdatesPaused  *bool      `db:"updates_paused"`
	PausedTime     *time.Time `db:"paused_time"`
	Reference      string     `db:"reference_text"`
}
