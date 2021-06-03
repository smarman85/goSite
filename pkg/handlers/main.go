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

  "goSite/pkg/jokes"
)

type joke1 struct {
        Joke string `json:"joke"`
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
}

func NotFound(w http.ResponseWriter, r *http.Request) { // a * before http.Request
        err := tpl.ExecuteTemplate(w, "404.gohtml", nil)
        if err != nil {
                log.Fatalln("template didn't execute: ", err)
        }
}

func JokeByID(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        j := &joke1 {
                Joke: jokes.Jokes[id],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(w, string(jm))
        //fmt.Fprintf(res, fmt.Sprintf("joke: %v", jokes.Jokes[id]))
}

func Joke(w http.ResponseWriter, r *http.Request) {
        jokeID := strconv.Itoa(rand.Intn(len(jokes.Jokes)))

        j := &joke1 {
                Joke: jokes.Jokes[jokeID],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(w, string(jm))
}
