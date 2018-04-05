package main

import (
	"Training/Day15/item/routers"
	"log"
	"net/http"
)

func main() {
	// Pindah ke router, buat router.go dan buat fungsi routing
	router := routers.InitRouters()

	//lalu buat konfigurasi server
	log.Fatal(http.ListenAndServe(":8000", router))
}
