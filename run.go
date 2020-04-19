package main

import (
  //"fmt"
  "net/http"
  "html/template"

  "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
  return
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
  tmpl.Execute(w, nil)
}

func main() {
  r := mux.NewRouter()
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  r.HandleFunc("/", mainhandler)
  r.HandleFunc("/products", handler)//.Methods("POST")
  r.HandleFunc("/articles", handler)//.Methods("GET")
  r.HandleFunc("/articles/{id}", handler)//.Methods("GET", "PUT")
  r.HandleFunc("/authors", handler)//.Queries("surname", "{surname}")

  http.Handle("/", r)
  http.ListenAndServe(":8089", r)
}


/*
https://stackoverflow.com/questions/49198742/golang-handle-static-folder
https://medium.com/@hugo.bjarred/rest-api-with-golang-and-mux-e934f581b8b5
https://www.gorillatoolkit.org/pkg/mux
https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
https://github.com/gorilla/mux
https://stackoverflow.com/questions/55996263/how-to-serve-static-files-for-all-pages-not-just-a-few
https://www.alexedwards.net/blog/serving-static-sites-with-go
*/
