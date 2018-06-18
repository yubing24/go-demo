// Minimum code to get "Hello, world!"
package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8000", nil)
}
