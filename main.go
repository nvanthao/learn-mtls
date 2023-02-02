package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {

	port := "8443"

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the way!")
	})

	certFile := "./certs/server.pem"
	keyFile := "./certs/server-key.pem"
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Println("Failed to load server cert and key:", err)
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MaxVersion:   tls.VersionTLS12,
	}

	srv := &http.Server{
		Addr:      ":" + port,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	fmt.Printf("Starting HTTPS server on port: %s ...", port)
	err = srv.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Failed to start HTTPS server:", err)
	}

}
