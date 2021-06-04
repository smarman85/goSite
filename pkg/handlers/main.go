package handlers

import (
  "github.com/gorilla/mux"
  "encoding/json"
  "math/rand"
  "strconv"
  "fmt"
  "net/http"

  "goSite/pkg/jokes"
)

type joke1 struct {
        Joke string `json:"joke"`
}

func JokeByID(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        j := &joke1 {
                Joke: jokes.Jokes[id],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(w, string(jm))
}

func Joke(w http.ResponseWriter, r *http.Request) {
        jokeID := strconv.Itoa(rand.Intn(len(jokes.Jokes)))

        j := &joke1 {
                Joke: jokes.Jokes[jokeID],
        }
        jm, _ := json.Marshal(j)
        fmt.Fprintf(w, string(jm))
}
