package apiservice

import "github.com/gorilla/mux"

type Registry interface {
	RegisterApi()
	MovieApi()
}

type Routes struct{
	Router *mux.Router
}
//----------------------------------------------
// Register Routes
//----------------------------------------------
func(r *Routes)RegisterRoutes(){
	r.Router=mux.NewRouter().StrictSlash(true)
}