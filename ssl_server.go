package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type SSLHandler struct{}

func (h SSLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func main() {
	const sslDir = "ssl-cert" + string(os.PathSeparator)
	const caCert = sslDir + "ca.pem"
	const certChainFile = sslDir + "cert_chain.crt"
	const keyFile = sslDir + "decrypted.ssl.key"

	var handler SSLHandler

	rootPEM, err := ioutil.ReadFile(caCert)
	if err != nil {
		panic("Failed to read certificate file: " + err.Error())
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("Failed to parse root certificate")
	}

	config := &tls.Config{RootCAs: roots}
	server := http.Server{
		Addr:      ":443",
		Handler:   handler,
		TLSConfig: config,
	}

	err = server.ListenAndServeTLS(certChainFile, keyFile)
	// err = server.ListenAndServe()
	if err != nil {
		panic("Server failure: " + err.Error())
	}
}
