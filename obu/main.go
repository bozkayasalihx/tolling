package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"

	"github.com/bozkayasalihx/paid_road/types"
)

var sendInternal = time.Second * 2

const wsEndpoint = "ws://localhost:3000/ws"

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64((rand.Intn(100) + 1))
	t := rand.Float64()
	return n + t
}

func genOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	OBUIDS := genOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(OBUIDS); i++ {
			obu := types.OBUData{
				ID:   OBUIDS[i],
				Long: genCoord(),
				Lat:  genCoord(),
			}
			if err := conn.WriteJSON(obu); err != nil {
				log.Printf("couldn't write to websocket %v", err)
				continue
			}
		}
		time.Sleep(sendInternal)
	}
}
