package main

import (
	"log"
	"net/http"
	"strconv"

	"onikur.com/text-to-img-api/api"

	"onikur.com/text-to-img-api/utils"

	"onikur.com/text-to-img-api/conf"
)

func main() {
	conf.Init()
	utils.Fonts().CacheFonts()

	mux := http.NewServeMux()
	mth := &api.MakeTextHandler{}
	mux.Handle("/-/", mth)
	mux.Handle("/api/text/", mth)
	mux.Handle("/api/fonts", &api.ListFontsHandler{})

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(conf.Get().Server.Port), mux))
}
