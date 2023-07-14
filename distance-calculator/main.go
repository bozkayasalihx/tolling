package main

type KafkaConsumer struct {
	Consumer DataConsumer
}

func main() {
	consumer := NewKafkaConsumer()
	consumer.ReadMsgsLoop()
}
