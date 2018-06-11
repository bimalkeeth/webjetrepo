package common

import mod "webjetapi/models"

type IHelper interface {
	GetNotInMovieCinema(movie *mod.MovieHeaderVM,movielist *[]mod.MovieHeaderVM)([]mod.MovieHeaderVM)
}
type Helper struct {}

func(h *Helper)GetNotInMovieCinema(parent *[]mod.MovieHeaderVM,child *[]mod.MovieHeaderVM)([]mod.MovieHeaderVM){

	retlist:=make([]mod.MovieHeaderVM,0)
	if len(*child)>0{
		 for _,flm :=range *child{
			 id := flm.ID[2:len(flm.ID)]
			 found:=false
			 for _,cin :=range *parent{
			 	 pid:=cin.ID[2:len(cin.ID)]
			 	 if id==pid{
					 found=true
					 break
				 }
			 }
             if !found{
				 flm.NotExists=true
				 retlist=append(retlist,flm)
				 found=false
			 }
		 }
	}
    return retlist
}