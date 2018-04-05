package main

import (
	"Training/Day15/officer/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()
	log.Fatal(http.ListenAndServe(":8001", router))
}
