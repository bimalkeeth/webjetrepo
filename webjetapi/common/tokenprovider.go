package common

import m "webjetapi/models"
import cl "github.com/sethgrid/pester"

type IToken interface {
	GetToken()(*m.TokenVM)
	GetAddress(key string)(string)
	GetHttpClient()(*cl.Client)
}

type Token struct{}

var addresses map[string]string

func init(){
	addresses=getAddresses()
}

func(t *Token) GetToken()(*m.TokenVM){

	token:=new(m.TokenVM)
	token.Key="x-access-token"
	token.Value="sjd1HfkjU83ksdsm3802k"
	return token
}

func getAddresses()(map[string]string){

	apitList:= make(map[string]string)
	apitList["CINEMA_01"]="http://webjetapitest.azurewebsites.net/api/cinemaworld/movies"
	apitList["CINEMA_02"]="http://webjetapitest.azurewebsites.net/api/cinemaworld/movie/%s"
	apitList["FILM_01"]="http://webjetapitest.azurewebsites.net/api/filmworld/movies"
	apitList["FILM_02"]="http://webjetapitest.azurewebsites.net/api/filmworld/movie/%s"

	return apitList
}

func(t *Token)GetAddress(key string)(string){
	return addresses[key]
}

func(t *Token)GetHttpClient()(*cl.Client){
	client := cl.New()
	client.Concurrency = 3
	client.MaxRetries = 5
	client.Backoff = cl.ExponentialBackoff
	client.KeepLog = true
	return client
}