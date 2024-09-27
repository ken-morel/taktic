package taktik

import "github.com/gorilla/mux"

var router *mux.Router

func GetRouter() *mux.Router {
	router = mux.NewRouter()
	router.HandleFunc("/", do_index)
	return router
}
