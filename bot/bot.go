package bot

import (
	"fmt"

	"github.com/Seew0/MeetTuber/checkers"
	"github.com/playwright-community/playwright-go"
)

func Bot() {

	browserSource := "chrome"
	source := "Invidious"
	colorscheme := "dark"
	TimezoneId := "Asia/Kolkata"
	locale := "en-US"
	pw, err := playwright.Run()
	checkers.CheckForError(err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless:          playwright.Bool(false),
		Channel:           &browserSource,
		IgnoreDefaultArgs: []string{"--disable-component-extensions-with-background-pages"},
		Args: []string{
			fmt.Sprintf("--auto-select-desktop-capture-source= %s", source),
			"--use-fake-device-for-media-stream",
			"--use-fake-ui-for-media-stream",
			"--start-maximized",
		},
	})
	checkers.CheckForError(err)
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		ColorScheme: (*playwright.ColorScheme)(&colorscheme),
		Viewport:    nil,
		Locale:      &locale,
		TimezoneId:  &TimezoneId,
	})
	checkers.CheckForError(err)
	page, err := browser.NewPage()
	checkers.CheckForError(err)
	session, err := context.NewCDPSession(page)
	checkers.CheckForError(err)
	// fmt.Print(session)
	windowID, err := session.Send("Browser.getWindowForTarget", nil)
	checkers.CheckForError(err)
	bounds := map[string]interface{}{}
	session.Send("Browser.setWindowBounds", map[string]interface{}{
		"windowId": windowID,
		"bounds":   bounds,
	})
	// defer browser.Close()
}
