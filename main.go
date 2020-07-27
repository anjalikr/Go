package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello World!!!\n")

	})
	http.ListenAndServe(":8080", nil)

}
