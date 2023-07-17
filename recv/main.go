package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/bozkayasalihx/paid_road/types"
)

type DataRecv struct {
	msgch    chan types.OBUData
	wsConn   *websocket.Conn
	Producer DataProducer
}

func NewDataRecv() (*DataRecv, error) {
	var (
		prod DataProducer
		err  error
	)
	prod, err = NewKafkaDataConsumer()
	if err != nil {
		return nil, err
	}

	prod = NewLoggingMiddleware(prod)

	return &DataRecv{
		msgch:    make(chan types.OBUData, 128),
		Producer: prod,
	}, nil
}

func main() {
	config := types.NewConfig()
	recv, err := NewDataRecv()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.wsHandler)

	fmt.Println("data recv working...")

	http.ListenAndServe(config.WSEndpoint, nil)
}

func (dr *DataRecv) produceData(d types.OBUData) error {
	return dr.Producer.ProduceData(d)
}

func (dr *DataRecv) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.wsConn = conn
	go dr.wsRecvLoop()
}

func (dr *DataRecv) wsRecvLoop() {
	fmt.Println("client connected")
	for {
		var data types.OBUData
		if err := dr.wsConn.ReadJSON(&data); err != nil {
			log.Println("got errr: ", err)
			continue
		}
		if err := dr.produceData(data); err != nil {
			log.Fatal(err)
		}
	}
}
