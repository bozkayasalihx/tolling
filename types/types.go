package types

import "time"

type Config struct {
	WSEndpoint    string
	Topic         string
	KafkaEndpoint string
}

func NewConfig() *Config {
	return &Config{
		WSEndpoint:    ":3000",
		Topic:         "test-topic",
		KafkaEndpoint: "localhost:9092",
	}
}

type Invoice struct {
	ID            int       `json:"id"`
	TotalDistance float64   `json:"totaldistance"`
	Unix          time.Time `json:"unix"`
}

type OBUData struct {
	ID   int     `json:"id"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
