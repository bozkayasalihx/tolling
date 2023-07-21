package main

import (
	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/types"
)

type LoggingMiddleware struct {
	next DataProducer
}

func NewLoggingMiddleware(p DataProducer) *LoggingMiddleware {
	return &LoggingMiddleware{
		next: p,
	}
}

func (l *LoggingMiddleware) ProduceData(d types.OBUData) error {
	logrus.WithFields(logrus.Fields{
		"OBUID": d.ID,
		"Lat":   d.Lat,
		"Long":  d.Long,
	}).Info("Produced OBU data")
	return l.next.ProduceData(d)
}
