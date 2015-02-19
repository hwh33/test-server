package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	sslDir             = "cert" + string(os.PathSeparator)
	certChainFile      = sslDir + "cert_chain.crt"
	keyFile            = sslDir + "decrypted.ssl.key"
	welcomeMessageFile = "WelcomeMessage.txt"
	logFileName        = "error_log.txt"
)

type SSLHandler struct {
	WelcomeMessage string
}

func (h SSLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.WelcomeMessage)
}

type FileWriter struct {
	OpenFile os.File
}

func (w FileWriter) Write(p []byte) (int, error) {
	return w.OpenFile.Write(p)
}

func main() {

	// Read the welcome message from file and use it to create the HTTP handler.
	welcomeMessageBytes, err := ioutil.ReadFile(welcomeMessageFile)
	handler := SSLHandler{
		WelcomeMessage: string(welcomeMessageBytes),
	}

	// If the log file doesn't exist, create it.
	// Otherwise, open it for appending.
	var logFile *os.File
	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		logFile, err = os.Create(logFileName)
	} else {
		logFile, err = os.OpenFile(logFileName, os.O_RDWR|os.O_APPEND, 0660)
	}
	if err != nil {
		fmt.Println("Error creating error log file: " + err.Error())
	}

	// Create an error log which will append reported errors to the log file.
	logFileWriter := FileWriter{
		OpenFile: *logFile,
	}
	errorLog := log.New(logFileWriter, "", log.Ldate|log.Ltime)

	// Initialize our server to listen on the HTTPS port.
	server := http.Server{
		Addr:     ":https",
		Handler:  handler,
		ErrorLog: errorLog,
	}

	// Start the server.
	err = server.ListenAndServeTLS(certChainFile, keyFile)
	if err != nil {
		panic("Server failure: " + err.Error())
	}

}
