package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

// Flags struct
type Flags struct {
	Address string
	Dynamic string
}

// Content of backend
type Content struct {
	Dynamic string `json:"dynamic"`
	Static  string `json:"static"`
}

func main() {
	f := Flags{}
	flag.StringVar(&f.Address, "address", ":9000", "address for backend service")
	flag.StringVar(&f.Dynamic, "dynamic_content", "", "dynamic content for backend")
	flag.Parse()

	log.Printf("configuration: %+v", f)

	http.HandleFunc("/healtcheck", func(w http.ResponseWriter, r *http.Request) {
		probe := r.FormValue("probe")

		if probe == "readiness" {
			log.Println("healthcheck: readiness")
			w.WriteHeader(http.StatusOK)
			return
		}

		if probe == "liveness" {
			log.Println("healthcheck: liveness")
			w.WriteHeader(http.StatusOK)
			return
		}
	})

	http.HandleFunc("/v1/content", func(w http.ResponseWriter, r *http.Request) {
		c := Content{
			Static:  "this is static content",
			Dynamic: f.Dynamic,
		}

		out, err := json.Marshal(c)
		if err != nil {
			c = Content{
				Static:  err.Error(),
				Dynamic: err.Error(),
			}
			out, _ = json.Marshal(c)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(out)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	})

	log.Printf("Backend: running on address %s", f.Address)
	http.ListenAndServe(f.Address, nil)
}
