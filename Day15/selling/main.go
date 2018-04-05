package main

import (
	"Training/Day15/selling/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()
	log.Fatal(http.ListenAndServe(":8002", router))
}
