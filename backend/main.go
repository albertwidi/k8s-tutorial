package main

import (
	"log"
	"net/http"
)

const address = ":9000"

func main() {
	http.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
		return
	})

	log.Printf("Backend: running on address %s", address)
	http.ListenAndServe(address, nil)
}
