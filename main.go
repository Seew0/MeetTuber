package main

import (
	"github.com/Seew0/MeetTuber/api"
)

func main(){
	Url := api.GetvideoData()
	api.Openbrowser(Url)
}