package main

import (
	"github.com/bozkayasalihx/paid_road/types"
	"go.uber.org/zap"
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("logging data",
		"obuID", d.ID,
		"Long", d.Long,
		"Lat", d.Lat,
	)
	return l.next.ProduceData(d)
}
