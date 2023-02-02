package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := "8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the way!"))
	})

	fmt.Printf("HTTP server listening on port %s ...", port)
	http.ListenAndServe(":"+port, nil)
}
