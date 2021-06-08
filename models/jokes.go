package models

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"goSite/pkg/jokes"
)

type Data struct {
	Joke   string `json:"joke"`
	ID     string `json:"id"`
	Status int    `json:status`
}

func NotFound(id string) []byte {
	j := &Data{
		Joke:   "This is not the joke you are looking for.",
		ID:     id,
		Status: 404,
	}
	jm, _ := json.Marshal(j)
	return jm
}

func GetJoke() []byte {
	jokeID := strconv.Itoa(rand.Intn(len(jokes.Jokes)))
	j := &Data{
		Joke:   jokes.Jokes[jokeID],
		ID:     jokeID,
		Status: 200,
	}
	jm, _ := json.Marshal(j)
	return jm
}

func ByID(id string) []byte {
	jokeID, _ := strconv.Atoi(id)
	if jokeID > len(jokes.Jokes) || jokeID < 0 {
		return NotFound(id)
	} else {
		j := &Data{
			Joke:   jokes.Jokes[id],
			ID:     id,
			Status: 200,
		}
		jm, _ := json.Marshal(j)
		return jm
	}
}
