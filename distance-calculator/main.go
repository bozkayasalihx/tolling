package main

func main() {

	calcSrv := NewCalculateService()
	kafkaConsumer := NewKafkaConsumer(calcSrv)
	calcSrv = NewConsumerLogMiddlware(calcSrv)

	kafkaConsumer.Start()
}
