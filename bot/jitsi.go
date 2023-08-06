package bot

import (
	"fmt"
	"time"

	"github.com/Seew0/MeetTuber/checkers"
)

func JoinJitsiMeet(url string) {
	context := NewBrowser()
	jitsi, err := context.NewPage()
	checkers.CheckForError(err)
	jitsi.Goto(url)

	jitsi.Click("#videoconference_page > div.premeeting-screen.css-1a5i9rv-container > div:nth-child(1) > div > div > label > div.css-1xhc3mw-activeArea > div > svg")
	jitsi.Click("#premeeting-name-input")
	jitsi.Keyboard().Type("MeetTuber")
	jitsi.Click("#videoconference_page > div.premeeting-screen.css-1a5i9rv-container > div:nth-child(1) > div > div > div.css-19c54gc-inputContainer > div.css-1kahyo4-dropdownContainer > div > div")
	time.Sleep(time.Second *2)
	jitsi.Keyboard().Press("M")
	jitsi.Keyboard().Press("V")

	// jitsi.Click("#new-toolbox > div > div > div > div:nth-child(3) > div > div > div > svg > path:nth-child(1)")
	
	for true{
		fmt.Print("")
	}
}
