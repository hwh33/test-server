/* This is essentially the simple HTTP server from the 'Tour of Go' tutorial.
 *
 * It can be found here: https://tour.golang.org/methods/13
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Read the welcome message from file and print it.
var welcome_msg_file = "WelcomeMessage.txt"
var welcome_msg_buffer, err = ioutil.ReadFile(welcome_msg_file)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, string(welcome_msg_buffer))
}

func main() {
	if err != nil {
		log.Fatal(err)
		return
	}

	var h Hello
	err := http.ListenAndServe(":80", h)
	if err != nil {
		log.Fatal(err)
	}
}
