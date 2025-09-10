package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand/v2"
	"net/http"
)

var (
	rtp float64
)

type Response struct {
	Result float64 `json:"result"`
}

func getHandler(w http.ResponseWriter, _ *http.Request) {
	mul := rtp / rand.Float64()
	mul = min(max(mul, 1.0), 10000.0)

	json.NewEncoder(w).Encode(Response{Result: mul})
}

func main() {
	rtpFlag := flag.Float64("rtp", 0.95, "Target RTP value")
	flag.Parse()
	rtp = *rtpFlag

	http.HandleFunc("/get", getHandler)

	if err := http.ListenAndServe(":64333", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
