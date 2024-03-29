package models

type Movie struct{
  Movies []MovieHeaderVM
}

type MovieHeaderVM struct{
	ID string `json:"ID"`
	Title string `json:"Title"`
	Type string `json:"Type"`
	Poster string `json:"Poster"`
	Year string `json:"Year"`
	Detail MovieDetailVM `json:",omitempty"`
	NotExists bool `json:",omitempty"`
}
type MovieDetailVM struct{
	Title string `json:"Title"`
	Year string `json:"Year"`
	Rated string `json:"Rated"`
	Released string `json:"Released"`
	Runtime string `json:"Runtime"`
	Genre string `json:"Genre"`
	Director string `json:"Director"`
	Writer string `json:"Writer"`
	Actors string `json:"Actors"`
	Plot string `json:"Plot"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Awards string `json:"Awards"`
	Poster string `json:"Poster"`
	Metascore string `json:"Metascore"`
	Rating string `json:"Rating"`
	Votes string `json:"Votes"`
	ID string `json:"ID"`
	Type string `json:"Type"`
	Price string `json:"Price"`
}
