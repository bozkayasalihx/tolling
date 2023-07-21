package main

import (
	"context"
	"log"
	"time"

	"github.com/goccy/go-json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/aggr"
	"github.com/bozkayasalihx/paid_road/types"
)

type KafkaDataConsumer struct {
	Reader       *kafka.Reader
	CalculateSrv CalculateServicer
	DistSrv      aggr.AggrServicer
	isRunning    bool
}

func NewKafkaConsumer(calcSrv CalculateServicer, aggrSrv aggr.AggrServicer) *KafkaDataConsumer {
	config := types.NewConfig()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaEndpoint},
		Topic:     config.Topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	return &KafkaDataConsumer{
		Reader:       r,
		CalculateSrv: calcSrv,
		DistSrv:      aggrSrv,
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
		distance, err := c.CalculateSrv.CalculateDistance(d)
		if err != nil {
			log.Fatalf("Couldn't calculate distance: %v", err)
		}

		dist := types.Distance{
			ID:       d.ID,
			Distance: distance,
			Unix:     time.Now().Unix(),
		}

		if err := c.DistSrv.Aggregate(dist); err != nil {
			log.Fatalf("Couldn't aggregate distance: %v", err)
		}

	}
}
