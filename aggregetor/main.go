package main

import (
	"flag"
	"net/http"
)

func main() {

	endpoint := flag.String("endpoint", ":8058", "the endpoint of aggregetor")
	flag.Parse()
	http.HandleFunc("aggregate", handleAggregetor)
	http.ListenAndServe(*endpoint, nil)

}

func handleAggregetor(w http.ResponseWriter, r *http.Request) {}
