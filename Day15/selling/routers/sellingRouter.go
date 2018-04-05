package routers

import (
	cont "Training/Day15/selling/controllers"

	"github.com/gorilla/mux"
)

func setSellingRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/selling", cont.GetSelling).Methods("GET")
	return router
}
