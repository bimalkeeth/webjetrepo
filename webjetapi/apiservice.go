package main

import (
	"fmt"
	"webjetapi/apiservice"
	"net/http"
	"github.com/gorilla/handlers"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println("Service is running")
	movies:=new(apiservice.Routes)
	movies.RegisterRoutes()
	movies.MovieApi()
	log.Fatal(http.ListenAndServe(":8081",handlers.CORS()( movies.Router)))
	fmt.Println("Service is running")
}
