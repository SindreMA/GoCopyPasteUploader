package ClipboardService

import (
	ClipboardHelper "main/Helpers/Clipboard"
	"main/Types"
	"time"

	"golang.design/x/clipboard"
)

var ClipboardChanged = make(chan Types.IClipboard)

var ClipboardData *Types.IClipboard


func RefreshClipboardData(first bool) {
	data := ClipboardHelper.GetClipboardData()

	if ClipboardData != nil {
		if *ClipboardData == *data {
			return
		}	
	}

	ClipboardData = data

	if !first && data != nil {
	println("Clipboard changed", data.Data)
		//ClipboardChanged <- *data
	}
	
}

func SetupRefresher() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(250 * time.Millisecond)

	firstRun := true

	go func() {
		for range ticker.C {
			RefreshClipboardData(firstRun)
			firstRun = false
		}
	}()
}