package main

import (
	"flag"
	"net/http"
)

func main() {
	flag.Parse()

	http.ListenAndServe(":9000", nil)
}
