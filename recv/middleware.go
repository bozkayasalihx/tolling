package main

import (
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
	return l.next.ProduceData(d)
}
