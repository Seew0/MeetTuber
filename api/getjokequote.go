package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Seew0/MeetTuber/checkers"
)

type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Id        int    `json:"id"`
}

var APIURL string = "https://official-joke-api.appspot.com/random_joke"

func Getrandomjoke() string {
	var quote string = "Hey i am MeetTuber how are you today?"
	var newJoke Joke

	response, err := http.Get(APIURL)
	checkers.CheckForError(err)

	data, err := ioutil.ReadAll(response.Body)
	checkers.CheckForError(err)
	valid := checkers.CheckValideJson(data)

	if !valid {
		quote = "Cant Think of anything else right now" + quote
	} else {
		json.Unmarshal(data,&newJoke)

		quote = newJoke.Setup + "\n" + newJoke.Punchline + "\n" + "\n"+ quote
	}

	return quote
}
