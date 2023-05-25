package commands

import "github.com/Seew0/MeetTuber/api"

func GetPopular(){
	var URL string = api.GetvideoData()
	api.Openbrowser(URL)
}