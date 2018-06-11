package apiservice

//--------------------------------------------------
// Register webapi
//--------------------------------------------------
func(r *Routes)MovieApi(){
	r.Router.HandleFunc("/api/movies/", GetAllMovies).Methods("GET")

}
