package main

import (
	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/types"
)

type ConsumerLogMiddleware struct {
	next DataConsumer
}

func NewConsumerLogMiddlware(c DataConsumer) *ConsumerLogMiddleware {
	return &ConsumerLogMiddleware{
		next: c,
	}
}

func (c *ConsumerLogMiddleware) ConsumeData(d types.OBUData) error {
	logrus.WithFields(logrus.Fields{
		"OBUID": d.ID,
		"Lat":   d.Lat,
		"Long":  d.Long,
	}).Info("Data Consumer ->")
	return c.next.ConsumeData(d)
}
