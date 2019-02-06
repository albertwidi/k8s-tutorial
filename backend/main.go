package main

import "net/http"

func main() {
	http.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
		return
	})
	http.ListenAndServe(":9000", nil)
}
