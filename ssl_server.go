package main

import (
	"fmt"
	"net/http"
	"os"
)

type SSLHandler struct{}

func (h SSLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func main() {
	const sslDir = "ssl-cert" + string(os.PathSeparator)
	const certChainFile = sslDir + "cert_chain.crt"
	const keyFile = sslDir + "decrypted.ssl.key"

	var handler SSLHandler

	server := http.Server{
		Addr:      ":http",
		Handler:   handler,
	}

	// err := server.ListenAndServeTLS(certChainFile, keyFile)
	err := server.ListenAndServe()
	if err != nil {
		panic("Server failure: " + err.Error())
	}
}
