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
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":8080", sm)

}
