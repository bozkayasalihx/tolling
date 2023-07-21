package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/types"
)

type ConsumerLogMiddlware struct {
	next CalculateServicer
}

func NewConsumerLogMiddlware(srv CalculateServicer) CalculateServicer {
	return &ConsumerLogMiddlware{
		next: srv,
	}
}

func (c *ConsumerLogMiddlware) CalculateDistance(d types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"dist": dist,
			"err":  err,
		}).Info("consume data")
	}(time.Now())

	dist, err = c.next.CalculateDistance(d)
	return
}
