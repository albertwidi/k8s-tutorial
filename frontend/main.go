package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

// Flags struct
type Flags struct {
	Address string
}

// Page struct
type Page struct {
	Title   string
	Content string
}

func main() {
	f := Flags{}
	flag.StringVar(&f.Address, "address", ":9090", "address for frontend service")
	flag.Parse()

	p := Page{
		Title:   "k8s-tutorial-frontend",
		Content: "haloha",
	}

	tmpl, _ := template.New("frontend-page").Parse(asset)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, p)
		return
	})

	log.Printf("Frontend: running on address %s", f.Address)
	http.ListenAndServe(f.Address, nil)
}

const asset = `
<html>
<head>
	<title> {{ .Title }}
</head>
<body>
	{{ .Content }}
</body>
</html>
`
