package busrules

import "webjetapi/models"
import (man "webjetapi/manager"
	        "log"
	    hlp "webjetapi/common"
	"strconv"
)

type IRules interface {

	GetAllMovieData()([]models.MovieHeaderVM)
	GetAllMovieDataWithoutDetail()([]models.MovieHeaderVM)
	GetMovieById(id string)(models.MovieDetailVM)
}
type Rules struct {}

func(r *Rules)GetAllMovieDataWithoutDetail()([]models.MovieHeaderVM){

	manager:= new(man.MovieService)
	cinemaData,err:= manager.GetAllCinemaWorldMovies()
	if err!=nil{
		log.Fatal(err)
	}
	filmData,err:= manager.GetAllFilmWorldMovies()
	if err!=nil{
		log.Fatal(err)
	}
	help:=new(hlp.Helper)
	if len(cinemaData)> 0{
		var finallist=make([]models.MovieHeaderVM,0)
		itemNotExists:=help.GetNotInMovieCinema(&cinemaData,&filmData)
		for _,item:=range itemNotExists{

			item.ID=item.ID[2:len(item.ID)]
			finallist=append(finallist,item)
		}
		for _,mov := range cinemaData{

			mov.ID=mov.ID[2:len(mov.ID)]
			finallist=append(finallist,mov)
		}
		return finallist
	}else{
		var finallist=make([]models.MovieHeaderVM,0)
		itemNotExists:=help.GetNotInMovieCinema(&filmData,&cinemaData)
		for _,item:=range itemNotExists{

			item.ID=item.ID[2:len(item.ID)]
			finallist=append(finallist,item)
		}
		for _,mov := range filmData{
			mov.ID=mov.ID[2:len(mov.ID)]
			finallist=append(finallist,mov)
		}
		return finallist

	}
}

func(r *Rules)GetMovieById(id string)(models.MovieDetailVM){
	manager:= new(man.MovieService)
	ciena,err:= manager.GetCinemaWorldById("cw"+id)
	if err!=nil{
		log.Fatal("Error in retriving with GetFilmWorldById")
	}
	film,err:= manager.GetFilmWorldById("fw"+id)
	if err!=nil{
		log.Fatal("Error in retriving with GetFilmWorldById")
	}
	fimval,_:=strconv.ParseFloat(film.Price,64)
	cinval,_:=strconv.ParseFloat(ciena.Price,64)

	if cinval<fimval && cinval>0  {
		return ciena
	}else{
		return film
	}
}

