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

//var projects []string ("one", "two")

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
  router.HandleFunc("/projects/{project}", project)
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
  projects := map[string][]string{
    "test": []string{"test header", "quick desc"},
    "test1": []string{"test1 header", "quick desc 1"},
    "test2": []string{"test2 header", "quick desc 2"},
    "test3": []string{"test3 header", "quick desc 3"},
    "test4": []string{"test4 header", "quick desc 4"},
    "test5": []string{"New Post", "something catchy"},
    "post1": []string{"Lorem Ipsume", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
  }
  err := tpl.ExecuteTemplate(res, "projects.gohtml", projects)
  if err != nil {
    log.Fatalln("template didn't execute: ", err)
  }
}

func project(res http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  //id := params["project"]
  err := tpl.ExecuteTemplate(res, "project.gohtml", params)
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
