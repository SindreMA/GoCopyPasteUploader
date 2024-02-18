package main

import (
	"context"
	"fmt"
	"time"

	"golang.design/x/clipboard"
)

func main() {

	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(500 * time.Millisecond)

    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
			fmt.Println(clipboard.Read(clipboard.FmtText))
			fmt.Println(clipboard.Read(clipboard.FmtImage))
        }
    }()

	var test = clipboard.Read(clipboard.FmtText)

	cbTextChannel := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range cbTextChannel {
		// print out clipboard data whenever it is changed
		println("clipboard changed text", string(data))
	}


	cbImageChannel := clipboard.Watch(context.TODO(), clipboard.FmtImage)
	for data := range cbImageChannel {
		// print out clipboard data whenever it is changed
		println("clipboard changed image", data)
	}

	fmt.Println(test)
}
