package main

import (
	"context"

	"github.com/goccy/go-json"
	"github.com/segmentio/kafka-go"

	"github.com/bozkayasalihx/paid_road/types"
)

type DataProducer interface {
	ProduceData(types.OBUData) error
}

type kafkaDataProducer struct {
	Conn *kafka.Conn
}

func NewKafkaDataConsumer() (DataProducer, error) {
	config := types.NewConfig()
	conn, err := kafka.DialLeader(
		context.Background(),
		"tcp",
		config.KafkaEndpoint,
		config.Topic,
		0,
	)
	if err != nil {
		return nil, err
	}
	return &kafkaDataProducer{
		Conn: conn,
	}, nil
}

func (k *kafkaDataProducer) ProduceData(d types.OBUData) error {
	bb, err := json.Marshal(d)
	if err != nil {
		return err
	}
	_, err = k.Conn.Write([]byte(bb))
	if err != nil {
		return err
	}
	return nil
}
