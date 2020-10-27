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

//SyncAdjustments
func (as AdjustmentService) SyncAdjustments() {

	log.Debug().Msg("sync adjustments")

	for {
		adjustments, err := as.getAdjustments()
		log.Debug().
			Int("#adustments", len(adjustments)).
			Msg("adjustments found")

		if err != nil {
			log.Err(err).Msg("cannot get adjustments")
		}

		if err == nil && len(adjustments) > 0 {
			for i := 0; i < len(adjustments); i++ {
				if err := as.sage.addAdjustment(adjustments[i]); err != nil {
					_ = as.deleteAdjustment(adjustments[i].Id)
				}
			}
		}

		time.Sleep(time.Duration(as.sleep) * time.Second)

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
