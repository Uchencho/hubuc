package main

import (
	"log"
	"net/http"

	"github.com/Uchencho/hubuc/internal/app"
)

const defaultAddress = "127.0.0.1:8000"

func main() {
	a := app.New()

	log.Println("starting server on port: ", defaultAddress)
	log.Fatal(http.ListenAndServe(defaultAddress, a.Handler()))
}
