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
}

func NewAdjustmentService() AdjustmentService {
	return AdjustmentService{NewSage(), config.Config.SleepPeriod}
}

func (as AdjustmentService) SyncAdjustments() {

	log.Debug().Msg("Sync adjustments")

	for {
		adjustments, err := as.getAdjustments()

		if err != nil {
			log.Err(err).Msg("Cannot get adjustments")
		} else if len(adjustments) == 0 {
			log.Debug().Msg("No outstanding adjustments")
		} else {
			if err == nil && len(adjustments) > 0 {
				for i := 0; i < len(adjustments); i++ {
					if err := as.sage.addAdjustment(adjustments[i]); err != nil {
						_ = as.deleteAdjustment(adjustments[i].Id)
					}
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
	dbase, err := db.GetDefaultPersistenceImpl()

	if err != nil {
		log.Error().Err(err)
		return err
	}

	return dbase.DeleteAdjustment(id)

}

func (as AdjustmentService) getAdjustments() ([]types.Adjustments, error) {
	dbase, err := db.GetDefaultPersistenceImpl()

	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	adj, err := dbase.GetAdjustments()

	if err != nil {
		return nil, fmt.Errorf("cannot read from the adjustments table")
	}

	return adj, nil

}
