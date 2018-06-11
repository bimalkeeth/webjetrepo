package manager

import (
	mod "webjetapi/models"
	"net/http"
	"log"
	"encoding/json"
	com "webjetapi/common"
	"fmt"
)
type IMovieService interface {

	GetAllCinemaWorldMovies()([]mod.MovieHeaderVM,error)
	GetAllFilmWorldMovies()([]mod.MovieHeaderVM,error)
	GetCinemaWorldById(id string)(mod.MovieDetailVM,error)
	GetFilmWorldById(id string)(mod.MovieDetailVM,error)
}

type MovieService struct {}


func(m *MovieService)GetAllCinemaWorldMovies()([]mod.MovieHeaderVM,error){

	tok:=com.Token{}
	req, err := http.NewRequest("GET", tok.GetAddress("CINEMA_01"), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil,err
	}

	token:= tok.GetToken()
	req.Header.Set(token.Key,token.Value)

	client := tok.GetHttpClient()
	resp,err:=client.Do(req)

	if err != nil {
		log.Fatal("Do: ", err)
		return nil,err
	}
	defer resp.Body.Close()
	var rec mod.Movie

	if err := json.NewDecoder(resp.Body).Decode(&rec); err != nil {
		log.Fatal(err)
	}
	return rec.Movies,nil
}
func(m *MovieService)GetAllFilmWorldMovies()([]mod.MovieHeaderVM,error){

	tok:=com.Token{}
	req, err := http.NewRequest("GET", tok.GetAddress("FILM_01"), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil,err
	}

	token:= tok.GetToken()
	req.Header.Set(token.Key,token.Value)

	client := tok.GetHttpClient()
	resp,err:=client.Do(req)

	if err != nil {
		log.Fatal("Do: ", err)
		return nil,err
	}
	defer resp.Body.Close()
	var rec mod.Movie
	if err := json.NewDecoder(resp.Body).Decode(&rec); err != nil {
		log.Println(err)
	}
	return rec.Movies,nil
}
func(m *MovieService)GetCinemaWorldById(id string)(mod.MovieDetailVM,error){

	tok:=com.Token{}
	req, err := http.NewRequest("GET", fmt.Sprintf(tok.GetAddress("CINEMA_02"),id), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return mod.MovieDetailVM{},err
	}

	token:= tok.GetToken()
	req.Header.Set(token.Key,token.Value)

	client := tok.GetHttpClient()
	resp,err:=client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return mod.MovieDetailVM{},err
	}
	defer resp.Body.Close()
	var records= mod.MovieDetailVM{}
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}
	return records,nil

}
func(m *MovieService)GetFilmWorldById(id string)(mod.MovieDetailVM,error){
	tok:=com.Token{}
	req, err := http.NewRequest("GET", fmt.Sprintf(tok.GetAddress("FILM_02"),id), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return mod.MovieDetailVM{},err
	}
	token:= tok.GetToken()
	req.Header.Set(token.Key,token.Value)

	client := tok.GetHttpClient()
	resp,err:=client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return mod.MovieDetailVM{},err
	}
	defer resp.Body.Close()
	var records= mod.MovieDetailVM{}
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}
	return records,nil

}