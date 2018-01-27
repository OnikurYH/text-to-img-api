package main

import (
	"log"
	"net/http"

	"onikur.com/text-to-img-api/api"
)

func main() {
	m := http.NewServeMux()

	c := api.Controller{Mux: m}
	c.Init()

	log.Fatal(http.ListenAndServe(":8080", m))
}
