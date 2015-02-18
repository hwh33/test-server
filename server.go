/* This is the simple HTTP server from the 'Tour of Go' tutorial.
 *
 * It can be found here: https://tour.golang.org/methods/13
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:80", h)
	if err != nil {
		log.Fatal(err)
	}
}
