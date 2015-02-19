package main

import (
	"fmt"
	"net/http"
	"os"
)

type SSLHandler struct {
	WelcomeMessage string
}

func (h SSLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.WelcomeMessage)
}

func main() {
	const sslDir = "ssl-cert" + string(os.PathSeparator)
	const certChainFile = sslDir + "cert_chain.crt"
	const keyFile = sslDir + "decrypted.ssl.key"

	handler := SSLHandler{
		WelcomeMessage: "Welcome",
	}

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: handler,
	}

	// err := server.ListenAndServeTLS(certChainFile, keyFile)
	err := server.ListenAndServe()
	if err != nil {
		panic("Server failure: " + err.Error())
	}
}
