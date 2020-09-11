package main

import (
	"fmt"
	"log"
	"net/http"
	"surface"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
    buf := surface.Surface()
    fmt.Fprintf(w, (&buf).String())
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

