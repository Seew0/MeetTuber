package main

import (
	"github.com/Seew0/MeetTuber/checkers"
	"github.com/Seew0/MeetTuber/cli"
	"github.com/playwright-community/playwright-go"
)

func init() {
	err := playwright.Install()
	checkers.CheckForError(err)
}
func main() {
	cli.Run()
}
