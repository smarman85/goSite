package controllers

import (
  //"goSite/views"
  "net/http"
  //"goSite/models"
  //"goStie/pkg/jokes"
  //"github.com/gorilla/mux"
)

//func NewApi(r *mux.Router) *Jokes{
//  return &Jokes{
//    Joke: views.NewJoke(
//      "api/joke",
//    ),
//  }
//}

type Jokes struct {
  Jokes string `json:"joke"`
}

//type Jokes struct {
//  Joke *views.Api
//}

//func Joke(w http.ResponseWriter, r *http.Request) {
//}

func Random(w http.ResponseWriter, r *http.Request) {
}
