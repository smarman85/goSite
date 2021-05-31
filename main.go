// ref https://www.alexedwards.net/blog/serving-static-sites-with-go
// https://stackoverflow.com/questions/55996263/how-to-serve-static-files-for-all-pages-not-just-a-few
package main

import (
  "net/http"

  "goSite/pkg/handlers"
  "goSite/controllers"
  //"goSite/rand"

  //"github.com/gorilla/csrf"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
  staticC := controllers.NewStatic()
  postsC := controllers.NewPosts(r)


  //router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  //router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
  // static pages
  r.Handle("/", staticC.Home).Methods("GET")
  r.Handle("/about", staticC.About).Methods("GET")
  //r.Handle("/resume", staticC.Resume)

  // posts routes
  //r.HandleFunc("/projects", handlers.Projects)
  r.HandleFunc("/projects", postsC.ShowAll).Methods("GET")
  r.HandleFunc("/projects/{project}", handlers.Project)
  r.HandleFunc("/api/joke/{id}", handlers.JokeByID)
  r.HandleFunc("/api/joke", handlers.Joke)
  r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

  assetHandler := http.FileServer(http.Dir("./assets/"))
  assetHandler = http.StripPrefix("/assets/", assetHandler)
  r.PathPrefix("/assets/").Handler(assetHandler)

  http.ListenAndServe(":8088", r)
}
