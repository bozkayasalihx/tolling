package main

type KafkaConsumer struct {
	Consumer DataConsumer
}

func main() {
	var consumer DataConsumer

	consumer = NewKafkaConsumer()
	consumer = NewConsumerLogMiddlware(consumer)
}
