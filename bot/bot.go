package bot

import (
	"fmt"

	"github.com/Seew0/MeetTuber/checkers"
	"github.com/playwright-community/playwright-go"
)

func NewBrowser() playwright.BrowserContext {
	browserSource := "chrome"
	colorscheme := "dark"
	TimezoneId := "Asia/Kolkata"
	locale := "en-US"
	useragent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4595.0 Safari/537.36"
	screenShareSource := "invidious"
	pw, err := playwright.Run()
	checkers.CheckForError(err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless:          playwright.Bool(false),
		Channel:           &browserSource,
		IgnoreDefaultArgs: []string{"--disable-component-extensions-with-background-pages"},
		Args: []string{
			fmt.Sprintf("--auto-select-desktop-capture-source=%s", screenShareSource),
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
		UserAgent:   &useragent,
		TimezoneId:  &TimezoneId,
	})
	checkers.CheckForError(err)

	return context
}

func MinimizeBrowser(context playwright.BrowserContext, page playwright.Page) {
	type bounds struct {
		windowState string
	}

	session, err := context.NewCDPSession(page)
	checkers.CheckForError(err)

	windowID, err := session.Send("Browser.getWindowForTarget", nil)
	checkers.CheckForError(err)

	session.Send("Browser.setWindowBounds", map[string]interface{}{
		"windowId": windowID,
		"bounds": bounds{
			windowState: "minimized",
		},
	})
}

func CloseBrowser(context playwright.BrowserContext) {
	context.Close()
}
