package main

import (
	"log"
	"net/http"

	"onikur.com/text-to-img-api/utils"

	"onikur.com/text-to-img-api/conf"

	"onikur.com/text-to-img-api/api"
)

func main() {
	conf.Init()
	utils.Fonts().CacheFonts()

	m := http.NewServeMux()

	c := api.Controller{Mux: m}
	c.Init()

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", m))
}
