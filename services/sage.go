package services

import (
	"github.com/rs/zerolog/log"
	"lw-adjustments/types"
)

type Sage struct{}

func NewSage() Sage {
	return Sage{}
}

func (sage Sage) addAdjustment(adjustment types.Adjustments) error {
	log.Info().
		Str("stock code", adjustment.StockCode).
		Str("type", adjustment.AdjustmentType).
		Float64("amount", adjustment.Amount).
		Msg("updating Sage...")

	return nil
}
