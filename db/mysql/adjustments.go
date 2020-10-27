package mysql

import (
	"lw-adjustments/types"
)

func (s *Connection) GetAdjustments() ([]types.Adjustments, error) {
	res := s.DB.Collection("adjustments").Find()

	var adjustments []types.Adjustments
	err := res.All(&adjustments)
	if err != nil {
		return nil, err
	}
	return adjustments, nil
}

func (s *Connection) GetAllAdjustments() ([]types.Adjustments, error) {
	var adjustments []types.Adjustments

	res := s.DB.Collection("adjustments").Find()
	err := res.All(&adjustments)
	return adjustments, err
}

func (s *Connection) DeleteAdjustment(id int) error {
	return s.DB.Collection("adjustments").Find("id", id).Delete()
}
