package main

import "github.com/bozkayasalihx/paid_road/types"

type ConsumerLogMiddlware struct {
	next CalculateServicer
}

func NewConsumerLogMiddlware(srv CalculateServicer) CalculateServicer {
	return &ConsumerLogMiddlware{
		next: srv,
	}
}

func (c *ConsumerLogMiddlware) CalculateDistance(d types.OBUData) (dist float64, err error) {
	return 0.0, nil
}
