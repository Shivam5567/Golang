package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



type Movie struct{
	ID string `json:"id"`
	Isbn string `json:""`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main(){

	



}
