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
  router := mux.NewRouter()
  staticC := controllers.NewStatic()


  /*
  b, err := rand.Bytes(32)
  if err != nil {
    panic(err)
  }
  csrfMw := csrf.Protect(b, csrf.Secure(false))
  */

  //router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  //router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
  router.Handle("/", staticC.Home).Methods("GET")
  router.Handle("/about", staticC.About).Methods("GET")
  router.HandleFunc("/projects", handlers.Projects)
  router.HandleFunc("/projects/{project}", handlers.Project)
  router.HandleFunc("/resume", handlers.Resume)
  router.HandleFunc("/api/joke/{id}", handlers.JokeByID)
  router.HandleFunc("/api/joke", handlers.Joke)
  router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

  http.ListenAndServe(":8088", router)
}
