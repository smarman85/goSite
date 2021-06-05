package models

import (
  "encoding/json"
  "math/rand"
  "strconv"

  "goSite/pkg/jokes"
)

type Data struct {
  Joke string `json:"joke"`
}

func GetJoke() []byte {
  jokeID := strconv.Itoa(rand.Intn(len(jokes.Jokes)))
  j := &Data{
    Joke: jokes.Jokes[jokeID],
  }
  jm, _ := json.Marshal(j)
  return jm
  //fmt.Fprintf(w, string(jm))
}

func ByID(id string) []byte {
  j := &Data{
    Joke: jokes.Jokes[id],
  }
  jm, _ := json.Marshal(j)
  return jm
}
