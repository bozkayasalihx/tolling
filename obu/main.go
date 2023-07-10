package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var sendInternal = time.Second * 2

type OBUData struct {
	ID   int     `json:"id"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

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
	for {
		for i := 0; i < len(OBUIDS); i++ {
			obu := OBUData{
				ID:   OBUIDS[i],
				Long: genCoord(),
				Lat:  genCoord(),
			}
			fmt.Printf("new obu data [%d] <<Lat :: %v :: Long :: %v>>\n", obu.ID, obu.Lat, obu.Long)
		}
		time.Sleep(sendInternal)
	}
}
