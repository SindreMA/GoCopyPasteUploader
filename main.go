package main

import (
	"fmt"
	ClipboardService "main/services"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println(("Starting the clipboard service"))
	ClipboardService.SetupRefresher()
			/*
            fmt.Println("Tick at", t)
			fmt.Println(clipboard.Read(clipboard.FmtText))
			fmt.Println(clipboard.Read(clipboard.FmtImage))
			*/
	WaitAndHandleShutdown()
}

func WaitAndHandleShutdown() {
	// Setting up signal handling for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-signals
	fmt.Println("Received signal:", sig)
	fmt.Println("Shutting down gracefully...")
}