package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Didalam Folder Router buat file itemRouter.go
	//Guna buat funcHandler <EndPoint>
	router = setItemRouters(router)
	return router
}

//Setelah beres semua, menuju ke main.go
