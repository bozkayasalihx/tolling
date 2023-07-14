package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/types"
)

type KafkaDataConsumer struct {
	Reader           *kafka.Reader
	CalculateService CalculateServicer
	isRunning        bool
}

func NewKafkaConsumer(srv CalculateServicer) *KafkaDataConsumer {
	config := types.NewConfig()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaEndpoint},
		Topic:     config.Topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	return &KafkaDataConsumer{
		Reader:           r,
		CalculateService: srv,
	}
}

func (c *KafkaDataConsumer) Start() {
	c.isRunning = true
	c.ReadMsgsLoop()
}

// NOTE; making new  other method compatiable;
func (c *KafkaDataConsumer) ReadMsgsLoop() {
	logrus.Info("Reading messsages")
	for c.isRunning {
		msg, err := c.Reader.ReadMessage(context.Background())
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Info("couldn't read msgs")
			continue
		}

		var d types.OBUData
		if err := json.Unmarshal(msg.Value, &d); err != nil {
			logrus.Error(err)
			break
		}
		// NOTE: consume data from  serialization
		dist, err := c.CalculateService.CalculateDistance(d)
		fmt.Println(dist)
	}
}
