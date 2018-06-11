package apiservice

import "net/http"
import (
	bus "webjetapi/busrules"
	"encoding/json"
	 mod "webjetapi/models"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	manager:= new(bus.Rules)
	movielist:=manager.GetAllMovieDataWithoutDetail()
	movlist:=make([]mod.MovieHeaderVM,0)
	for _,mov :=range movielist{
		results:=GetMovieDetail(mov)
		for ch :=range results{
			movlist=append(movlist,ch)
		}
	}
	json.NewEncoder(w).Encode(movlist)
}
func GetMovieDetail(moviefrom mod.MovieHeaderVM)<-chan mod.MovieHeaderVM{
	chanMovie:= make(chan	mod.MovieHeaderVM)
	manager:= new(bus.Rules)
	go func( rule *bus.Rules,movie mod.MovieHeaderVM){
		detail:=rule.GetMovieById(movie.ID)
		movie.Detail=detail
		chanMovie <- movie
		close(chanMovie)
	}(manager,moviefrom)
	return chanMovie
}