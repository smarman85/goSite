// ref https://www.alexedwards.net/blog/serving-static-sites-with-go
// https://stackoverflow.com/questions/55996263/how-to-serve-static-files-for-all-pages-not-just-a-few
package main

import (
  //"fmt"
  "log"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
}

func main() {
  router:= mux.NewRouter()

  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
  router.HandleFunc("/", index)
  router.HandleFunc("/about", about)
  router.HandleFunc("/projects", projects)
  router.HandleFunc("/resume", resume)

  http.ListenAndServe(":8088", router)
}

func index(res http.ResponseWriter, req *http.Request) {
  err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
  if err != nil {
    log.Fatalln("template didn't execute: ", err)
  }
}

func about(res http.ResponseWriter, req *http.Request) {
  err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
  if err != nil {
    log.Fatalln("template didn't execute: ", err)
  }
}

func projects(res http.ResponseWriter, req *http.Request) {
  err := tpl.ExecuteTemplate(res, "projects.gohtml", nil)
  if err != nil {
    log.Fatalln("template didn't execute: ", err)
  }
}

func resume(res http.ResponseWriter, req *http.Request) {
  err := tpl.ExecuteTemplate(res, "resume.gohtml", nil)
  if err != nil {
    log.Fatalln("template didn't execute: ", err)
  }
}
