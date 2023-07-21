package aggr

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/types"
)

type AggrServicer interface {
	Aggregate(types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
}

type DistanceAggregator struct {
	store Storer
}

func NewDistanceAggr(s Storer) *DistanceAggregator {
	return &DistanceAggregator{
		store: s,
	}
}

func (d *DistanceAggregator) Aggregate(data types.Distance) error {
	defer func(s time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(s),
			"id":   data.ID,
			"dist": data.Distance,
		}).Info("aggregate data")
	}(time.Now())
	return d.store.Insert(data)
}
