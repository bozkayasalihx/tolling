package main

import (
	"errors"
	"math"

	"github.com/bozkayasalihx/paid_road/types"
)

type CalculateServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculateService struct {
	Points []float64
}

func NewCalculateService() CalculateServicer {
	return &CalculateService{
		Points: make([]float64, 0),
	}
}

func (s *CalculateService) CalculateDistance(data types.OBUData) (float64, error) {
	var (
		distance float64
		err      error
	)
	if len(s.Points) > 1 {
		distance, err = calcDistance([]float64{data.Long, data.Lat, s.Points[0], s.Points[1]})
	}
	s.Points = []float64{data.Long, data.Lat}
	return distance, err
}

func calcDistance(Coords []float64) (float64, error) {
	if len(Coords) > 4 {
		return 0.0, errors.New("too many arguments")
	}
	/// data {
	//  Long, Lat, Long, Lat
	//}
	d := math.Sqrt(math.Pow(Coords[2]-Coords[0], 2) + math.Pow(Coords[3]-Coords[1], 2))
	return d, nil
}
