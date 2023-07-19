package main

func main() {
	var (
		calcSrv CalculateServicer
	)

	calcSrv = NewCalculateService()
	calcSrv = NewConsumerLogMiddlware(calcSrv)
	kafkaConsumer := NewKafkaConsumer(calcSrv)

	kafkaConsumer.Start()
}
