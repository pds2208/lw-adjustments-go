package services

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"lw-adjustments/config"
	"lw-adjustments/db"
	"lw-adjustments/types"
	"time"
)

type AdjustmentService struct {
	sage  Sage
	sleep int
	db    db.Persistence
}

func NewAdjustmentService(db db.Persistence) AdjustmentService {
	return AdjustmentService{NewSage(), config.Config.SleepPeriod, db}
}

func (as AdjustmentService) SyncAdjustments() {

	log.Debug().Msg("Sync adjustments")

	for {
		if adjustments := as.getAdjustments(); adjustments != nil {
			log.Debug().Msg(fmt.Sprintf("%v adjustments to process", len(adjustments)))
			for _, adjustment := range adjustments {
				if err := as.sage.addAdjustment(adjustment); err != nil {
					_ = as.deleteAdjustment(adjustment.Id)
				}
			}
		}
		time.Sleep(time.Duration(as.sleep) * time.Second)
	}

}

func (as AdjustmentService) addAdjustment(adjustment types.Adjustments) {

	// Firstly we get the current sage quantity
	//qty, error := as.sage.getProductQuantity(adjustment.StockCode)

	if err := as.sage.addAdjustment(adjustment); err != nil {
		_ = as.deleteAdjustment(adjustment.Id)
	}

}

func (as AdjustmentService) deleteAdjustment(id int) error {
	return as.db.DeleteAdjustment(id)
}

func (as AdjustmentService) getAdjustments() []types.Adjustments {
	adj, err := as.db.GetAdjustments()
	if err != nil {
		log.Err(err).Msg("cannot read from the adjustments table")
		return nil
	}

	if len(adj) == 0 {
		log.Debug().Msg("No outstanding adjustments")
		return nil
	}

	return adj

}
