package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
)

func main() {
	port := "8443"

	// load server cert and key
	cert, err := tls.LoadX509KeyPair("./certs/server.pem", "./certs/server-key.pem")
	if err != nil {
		panic("Failed to load server cert and key: " + err.Error())
	}

	// load CA certificate that signed client cert
	caCert, err := tls.LoadX509KeyPair("./certs/ca.pem", "./certs/ca-key.pem")
	if err != nil {
		panic("Failed to load CA cert and key: " + err.Error())
	}
	x509CAcert, err := x509.ParseCertificate(caCert.Certificate[0])
	if err != nil {
		panic("Failed to parse CA cert: " + err.Error())
	}

	// create certificate pool
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(x509CAcert)

	// configure server's TLS settings
	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		MaxVersion:   tls.VersionTLS12,
	}

	// create HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the way!")
	})

	srv := &http.Server{
		Addr:      ":" + port,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	fmt.Printf("Starting HTTPS server with mTLS on port: %s ...", port)
	err = srv.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Failed to start HTTPS server:", err)
	}

}
