package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":5000", playerServer); err != nil {
		log.Fatalf("cound not listen on port 5000 %v", err)
	}
}
