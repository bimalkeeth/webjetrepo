import {Inject, Injectable} from '@angular/core';
import {DataService} from "./data.service";
import {IMovie} from "../models/movie.vm";
import {IMovieDetal as MovieDetal} from "../models/moviedetail.vm";
import {Observable} from "rxjs/Observable";
@Injectable()
export class BusinessrulesService {

  constructor(@Inject(DataService) public dataservice:DataService)
  {

  }
  GetAllMovies():Observable<IMovie[]>{
    return this.dataservice.GetAllCinemaServiceData()
  }
}



