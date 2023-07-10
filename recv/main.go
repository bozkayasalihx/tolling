package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/bozkayasalihx/paid_road/types"
)

const wsEndpoint = ":3000"

type DataRecv struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
}

func NewDataRecv() *DataRecv {
	return &DataRecv{
		msgch: make(chan types.OBUData, 128),
	}
}

func main() {
	recv := NewDataRecv()
	http.HandleFunc("/ws", recv.wsHandler)
	fmt.Println("data recv working...")

	http.ListenAndServe(wsEndpoint, nil)
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
	dr.conn = conn
	go dr.wsRecvLoop()
}

func (dr *DataRecv) wsRecvLoop() {
	fmt.Println("client connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("got errr: ", err)
			continue
		}
		fmt.Printf("new obu data [%d] <<Lat :: %v :: Long :: %v>>\n", data.ID, data.Lat, data.Long)
		dr.msgch <- data
	}
}
