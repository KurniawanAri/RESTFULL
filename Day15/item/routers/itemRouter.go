package routers

import (
	cont "Training/Day15/item/controllers"

	"github.com/gorilla/mux"
)

func setItemRouters(router *mux.Router) *mux.Router {
	// Sebelum membuat HandleFunc buat fungsinya dahulu di Folder cntroller file itemController.go
	router.HandleFunc("/item", cont.GetItem).Methods("GET")
	return router
}

//Setelah selesai itemRouter lalu selesaikan file router.go
