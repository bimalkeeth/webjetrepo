import {BrowserModule } from '@angular/platform-browser';
import {NgModule } from '@angular/core';
import {AppComponent } from './app.component';
import {FormsModule} from "@angular/forms";
import {HomeComponent } from './home/home.component';
import {HttpClientModule} from "@angular/common/http";
import {CommonModule} from "@angular/common";
import {BusinessrulesService} from "./services/businessrules.service";
import {DataService} from "./services/data.service";
import {NgxSpinnerModule} from "ngx-spinner";


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
   ],
  imports: [
    BrowserModule,
    CommonModule,
    HttpClientModule,
    FormsModule,
    NgxSpinnerModule
   ],
  providers: [BusinessrulesService,DataService],
  bootstrap: [AppComponent]
})
export class AppModule  {


}
