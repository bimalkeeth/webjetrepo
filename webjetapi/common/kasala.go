package common

//import (
//	"strconv"
//	"webjetapi/models"
//)
//
//func(r *Rules)GetAllMovieData()([]models.MovieHeaderVM){
//
//	manager:= new(man.MovieService)
//	cinemaData,err:= manager.GetAllCinemaWorldMovies()
//	if err!=nil{
//		log.Fatal(err)
//	}
//	filmData,err:= manager.GetAllFilmWorldMovies()
//	if err!=nil{
//		log.Fatal(err)
//	}
//	help:=new(hlp.Helper)
//	if len(cinemaData)> 0{
//		var finallist=make([]models.MovieHeaderVM,0)
//		itemNotExists:=help.GetNotInMovieCinema(&cinemaData,&filmData)
//		for _,item:=range itemNotExists{
//			film,err:= manager.GetFilmWorldById(item.ID)
//			if err!=nil{
//				log.Fatal("Error in retriving with GetFilmWorldById")
//			}
//			item.Detail=film
//			finallist=append(finallist,item)
//		}
//		for _,mov := range cinemaData{
//
//			film,err:= manager.GetFilmWorldById("fw"+mov.ID[2:len(mov.ID)])
//			if err!=nil{
//				log.Fatal("Error in retriving with GetFilmWorldById")
//			}
//
//			ciena,err:= manager.GetCinemaWorldById(mov.ID)
//			if err!=nil{
//				log.Fatal("Error in retriving with GetFilmWorldById")
//			}
//			fimval,_:=strconv.ParseFloat(film.Price,64)
//			cinval,_:=strconv.ParseFloat(ciena.Price,64)
//
//			if fimval< cinval && fimval>0 {
//				filmheader:=models.MovieHeaderVM{}
//				filmheader.ID=film.ID
//				filmheader.Title=film.Title
//				filmheader.Poster=film.Poster
//				filmheader.Type=film.Type
//				filmheader.Year=film.Year
//				filmheader.Detail=film
//				finallist=append(finallist,filmheader)
//			}else{
//				mov.Detail=ciena
//				finallist=append(finallist,mov)
//			}
//		}
//		return finallist
//	}else{
//		var finallist=make([]models.MovieHeaderVM,0)
//		itemNotExists:=help.GetNotInMovieCinema(&filmData,&cinemaData)
//		for _,item:=range itemNotExists{
//			ciena,err:= manager.GetCinemaWorldById(item.ID)
//			if err!=nil{
//				log.Fatal("Error in retriving with GetFilmWorldById")
//			}
//			item.Detail=ciena
//			finallist=append(finallist,item)
//		}
//		for _,mov := range filmData{
//
//			ciena,err:= manager.GetCinemaWorldById("cw"+mov.ID[2:len(mov.ID)])
//			if err!=nil{
//				log.Fatal("Error in retriving with GetFilmWorldById")
//			}
//
//			film,err:= manager.GetFilmWorldById(mov.ID)
//			if err!=nil{
//				log.Fatal("Error in retriving with GetFilmWorldById")
//			}
//			fimval,_:=strconv.ParseFloat(film.Price,64)
//			cinval,_:=strconv.ParseFloat(ciena.Price,64)
//
//			if cinval<fimval && cinval>0  {
//				filmheader:=models.MovieHeaderVM{}
//				filmheader.ID=ciena.ID
//				filmheader.Title=ciena.Title
//				filmheader.Poster=ciena.Poster
//				filmheader.Type=ciena.Type
//				filmheader.Year=ciena.Year
//				filmheader.Detail=ciena
//				finallist=append(finallist,filmheader)
//			}else{
//				mov.Detail=film
//				finallist=append(finallist,mov)
//			}
//		}
//		return finallist
//
//	}
//}
