import {Inject, Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {IMovie} from "../models/movie.vm";
import {Observable} from "rxjs/Observable";
import "rxjs/Rx"
import {IMovieDetal as MovieDetal} from "../models/moviedetail.vm";
import {MovieRange} from "../models/movierange.vm";
import {tap} from "rxjs/operators";

@Injectable()
export class DataService {

  constructor(@Inject(HttpClient) public http:HttpClient ) { }
  public options = {withCredentials:true};

  GetAllCinemaServiceData():Observable<IMovie[]>{

    let result:IMovie[]=[] ;
    return this.http.get<IMovie[]>("http://localhost:8081/api/movies/")
      .pipe( tap(data => console.log(JSON.stringify(data))), )

   }
}
