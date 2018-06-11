import {Component, Inject, OnInit} from '@angular/core';
import {BusinessrulesService} from "./services/businessrules.service";
import {IMovie} from "./models/movie.vm";
import {NgxSpinnerService} from "ngx-spinner";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'app';
  movieCollection:IMovie[]=[] ;
  constructor(@Inject(BusinessrulesService) public busservice:BusinessrulesService,private spinner: NgxSpinnerService)
  {

  }
  ngOnInit(): void {
    this.spinner.show();
    this.busservice.GetAllMovies().subscribe(s=>{
      this.movieCollection=s;
      this.spinner.hide();
    });
  }
}
