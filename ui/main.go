package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"ui/core"
	"ui/layouts"
	"ui/state"
	"ui/views"
	"usage"

	xenaC2 "github.com/zarkones/xena-client"
	"golang.design/x/clipboard"

	"fyne.io/fyne/v2"
)

func failedLicenseCallback() {
	views.Alert("Invalid license!")
}

func main() {
	usage.NewEvent(usage.APP_STARTED, "", nil)
	defer func() {
		if err := recover(); err != nil {
			usage.NewEvent(usage.PANIC, fmt.Sprint(err), nil)
			return
		}
		usage.NewEvent(usage.APP_STOPPED, "", nil)
	}()

	flag.Parse()

	// TODO: Run this code if it's supported platform.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	// state.AuthToken = os.Getenv("AUTH_TOKEN")
	// if state.AuthToken == "" {
	// 	views.Alert("AUTH_TOKEN environment variable not set. C2 will refuse to authorize user actions.")
	// 	usage.NewEvent(usage.AUTH_TOKEN_UNSET, "", nil)
	// }
	state.AuthToken = os.Getenv("AUTH_TOKEN")
	xenaC2.AuthToken = &state.AuthToken
	xenaC2.Init(&state.C2Host, nil, time.Minute*5)

	core.App.Settings().SetTheme(&mainTheme{})
	core.MainW.Resize(fyne.NewSize(core.WIN_WIDTH, core.WIN_HEIGHT))
	core.MainW.CenterOnScreen()

	core.MainW.SetContent(layouts.Default())

	core.MainW.ShowAndRun()
	os.Exit(0)
}
