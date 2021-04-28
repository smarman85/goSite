package handlers

import (
  "html/template"
  "github.com/gorilla/mux"
  "encoding/json"
  "math/rand"
  "strconv"
  "log"
  "fmt"
  "net/http"

  "goSite/pkg/posts"
  "goSite/pkg/jokes"
)

type joke1 struct {
        Joke string `json:"joke"`
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
}

func NotFound(res http.ResponseWriter, req *http.Request) { // a * before http.Request
        err := tpl.ExecuteTemplate(res, "404.gohtml", nil)
        if err != nil {
                log.Fatalln("template didn't execute: ", err)
        }
}

func Index(res http.ResponseWriter, req *http.Request) {
        err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
        if err != nil {
                log.Fatalln("template didn't execute: ", err)
        }
}

func About(res http.ResponseWriter, req *http.Request) {
        err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
        if err != nil {
                log.Println("template didn't execute: ", err)
                NotFound(res, req)
        }
}

func Projects(res http.ResponseWriter, req *http.Request) {
        projects := posts.Projects
        err := tpl.ExecuteTemplate(res, "projects.gohtml", projects)
        if err != nil {
                log.Println("template didn't execute: ", err)
                NotFound(res, req)
        }
}

func Project(res http.ResponseWriter, req *http.Request) {
        params := mux.Vars(req)
        //id := params["project"]
        err := tpl.ExecuteTemplate(res, params["project"]+".gohtml", nil)
        if err != nil {
                log.Println("template didn't execute: ", err)
                NotFound(res, req)
                //tpl.ExecuteTemplate(res, "404.gohtml", nil)
        }
}

func Resume(res http.ResponseWriter, req *http.Request) {
        err := tpl.ExecuteTemplate(res, "resume.gohtml", nil)
        if err != nil {
                log.Println("template didn't execute: ", err)
                NotFound(res, req)
        }
}

func JokeByID(res http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]

        j := &joke1 {
                Joke: jokes.Jokes[id],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(res, string(jm))
        //fmt.Fprintf(res, fmt.Sprintf("joke: %v", jokes.Jokes[id]))
}

func Joke(res http.ResponseWriter, req *http.Request) {
        jokeID := strconv.Itoa(rand.Intn(len(jokes.Jokes)))

        j := &joke1 {
                Joke: jokes.Jokes[jokeID],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(res, string(jm))
}
