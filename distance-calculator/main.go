package main

import "github.com/bozkayasalihx/paid_road/aggr"

func main() {
	var (
		calcSrv CalculateServicer
		aggrSrv aggr.AggrServicer
	)

	store := aggr.NewInMemStore()

	calcSrv = NewCalculateService()
	aggrSrv = aggr.NewDistanceAggr(store)
	calcSrv = NewConsumerLogMiddlware(calcSrv)
	kafkaConsumer := NewKafkaConsumer(calcSrv, aggrSrv)

	kafkaConsumer.Start()
}
