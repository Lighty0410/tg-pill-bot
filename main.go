package main

import (
	"log"
	"tg-pill-bot/server"
)

func main() {
	h := server.NewHttpServer()
	err := h.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
