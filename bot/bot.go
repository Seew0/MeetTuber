package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/Seew0/MeetTuber/checkers"
	"github.com/tebeka/selenium"
)

func Bot() {
	const (
		seleniumPath    = "drivers/selenium-server.jar"
		geckoDriverPath = "drivers/geckodriver"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(), //start x frame buffer for browser to run in
		selenium.GeckoDriver(geckoDriverPath),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)

	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	checkers.CheckForError(err)
	defer service.Stop()
	//web driver and connecting to local web browser

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	checkers.CheckForError(err)
	defer wd.Quit()

	if err := wd.Get("https://meet.jit.si/gobot"); err != nil {
		log.Fatal(err)
	}
}
