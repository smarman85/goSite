package controllers

import (
  "net/http"
  //"strconv"

  "goSite/models"
  "goSite/views"

  //"goStie/pkg/jokes"
  "github.com/gorilla/mux"
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
  joke := models.GetJoke()
  views.ApiRender(w, joke)
}

func Joke(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  //jokeID, err := strconv.Atoi(vars["id"])
  jokeID := vars["id"]
  joke := models.ByID(jokeID)
  views.ApiRender(w, joke)
}
