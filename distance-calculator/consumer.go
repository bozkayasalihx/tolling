package main

import (
	"context"
	"log"

	"github.com/bozkayasalihx/paid_road/types"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaConsumer struct {
	conn      *kafka.Conn
	isRunning bool
}

func NewKafkaConsumer() (*KafkaConsumer, error) {
	config := types.NewConfig()
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.KafkaEndpoint, config.Topic, 0)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	return &KafkaConsumer{
		conn: conn,
	}, nil
}

func (c *KafkaConsumer) Start() {
	c.isRunning = true
	c.ReadMessagesLoop()
}

func (c *KafkaConsumer) ReadMessagesLoop() {
	b := make([]byte, 10e3) // 10KB max per message
	for c.isRunning {
		_, err := c.conn.Read(b)
		if err != nil {
			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatal(err)
			}
			logger.Info("unable to read msgs")
			continue
		}
	}

}

func main() {
	kConsumer, err := NewKafkaConsumer()
	if err != nil {
		log.Fatal(err)
	}
	kConsumer.Start()
}
