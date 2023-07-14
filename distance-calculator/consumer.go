package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/types"
)

type DataConsumer interface {
	ConsumeData(types.OBUData) error
}

type KafkaDataConsumer struct {
	Reader    *kafka.Reader
	Consumer  DataConsumer
	isRunning bool
}

func NewKafkaConsumer() DataConsumer {
	config := types.NewConfig()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaEndpoint},
		Topic:     config.Topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	return &KafkaDataConsumer{
		Reader: r,
	}
}

func (c *KafkaDataConsumer) ConsumeData(d types.OBUData) error {
	return c.Consumer.ConsumeData(d)
}

// NOTE; making new  other method compatiable;
func (c *KafkaDataConsumer) ReadMsgsLoop() {
	logrus.Info("Reading messsages")
	for {
		msg, err := c.Reader.ReadMessage(context.Background())
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Info("couldn't read msgs")
			continue
		}
		var d types.OBUData
		if err := json.Unmarshal(msg.Value, &d); err != nil {
			logrus.Fatalf("couldn't parse msgs: %v", err)
			break
		}
		if err = c.ConsumeData(d); err != nil {
			logrus.Fatal(err)
			break
		}
	}
}
