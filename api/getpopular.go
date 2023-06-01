package api

import (
	"math/rand"

	"git.sr.ht/~greenfoo/govidious"
	"github.com/Seew0/MeetTuber/checkers"
)

func GetPopularvideoData() string{
	g := govidious.New("https://inv.vern.cc",nil)

	response,err := g.Popular()
	checkers.CheckForError(err)

	randomVariable := rand.Intn(40)
	var VideoId string = response[randomVariable].VideoId

	var URL string = "https://inv.vern.cc/watch?v=" + VideoId	

	return URL
}

