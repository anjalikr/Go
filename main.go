package main

import (
	"log"
	"net/http"
	"os"

	"example.com/hello/Microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	//Create Handlers
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":8080", sm)

}
