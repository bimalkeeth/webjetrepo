import {Component, Inject, OnInit} from '@angular/core';
import {BusinessrulesService} from "../services/businessrules.service";
import {IMovie} from "../models/movie.vm";
import {IMovieDetal as MovieDetal} from "../models/moviedetail.vm";


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  movieCollection:IMovie[]=[] ;
  movieDetail:MovieDetal;
  constructor(@Inject(BusinessrulesService) public busservice:BusinessrulesService)
  {

  }
  ngOnInit()
  {
    let dd="dsvdasgjhg";
    this.busservice.GetAllMovies().subscribe(s=>{
      this.movieCollection=s
    });
    let Spee=""
  }

}
