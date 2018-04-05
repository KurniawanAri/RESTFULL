package routers

import (
	cont "Training/Day15/officer/controller"

	"github.com/gorilla/mux"
)

func setOfficeRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/office", cont.GetOffice).Methods("GET")
	return router
}
