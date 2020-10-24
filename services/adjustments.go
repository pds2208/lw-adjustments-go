package services

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"lw-adjustments/db"
	"lw-adjustments/types"
	"time"
)

type AdjustmentService struct {
}

func init() {

}

//SyncAdjustments
func SyncAdjustments() {

	log.Debug().
		Msg("sync adjustments")
	for {
		adjustments, err := getAdjustments()
		log.Debug().
			Int("#adustments", len(adjustments)).
			Msg("adjustments found")
		if err != nil {
			log.Err(err).
				Msg("cannot get attributes")
		}
		if err == nil && len(adjustments) > 0 {
			for i := 0; i < len(adjustments); i++ {
				if err := updateSage(adjustments[i]); err != nil {
					_ = deleteAdjustment(adjustments[i].Id)
				}
			}
		}

		time.Sleep(10 * time.Second)

	}

}

func updateSage(adjustment types.Adjustments) error {
	log.Info().
		Str("adjustment", adjustment.StockCode).
		Msg("updating Sage")
	return nil
}

func deleteAdjustment(id int) error {
	dbase, err := db.GetDefaultPersistenceImpl()
	if err != nil {
		log.Error().Err(err)
		return err
	}

	return dbase.DeleteAdjustment(id)

}

func getAdjustments() ([]types.Adjustments, error) {
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
