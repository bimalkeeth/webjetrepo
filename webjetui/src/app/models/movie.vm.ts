import {SafeStyle} from "@angular/platform-browser";
import {IMovieDetal} from "./moviedetail.vm";

export interface IMovie {
  ID:string;
  Title:string;
  Year:number;
  Type:string;
  Poster:string;
  Img:SafeStyle;
  Detail:IMovieDetal
}
