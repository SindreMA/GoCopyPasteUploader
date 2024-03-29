package main

import (
	"flag"
	"fmt"
	ClipboardHelper "main/Helpers/Clipboard"
	FtpHelper "main/Helpers/Ftp"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"golang.design/x/clipboard"
)

func main() {
	//fmt.Println(("Starting the clipboard service"))
	
	//ClipboardService.SetupRefresher()

	// Define command-line flags
	server := flag.String("server", "", "The address of the FTP server")
	username := flag.String("username", "", "The username for authentication")
	password := flag.String("password", "", "The password for authentication")
	port := flag.Int("port", 21, "The port on which the FTP server is running (default is 21)")
	folder := flag.String("folder", "", "The folder to upload the file to")
	webServer := flag.String("webServer", "", "The web server the files will be hosted on")
	fileName := *flag.String("fileName", "", "The name of the file")

	// Parse the flags
	flag.Parse()

	// Simple validation to ensure required flags are provided
	if *server == "" || *username == "" || *password == "" {
		fmt.Println("server, username, and password flags are required")
		flag.PrintDefaults() // Print usage information
		return
	}
	
	data := ClipboardHelper.GetClipboardData()

	if data != nil {
		if &fileName == nil || fileName == "" {
			fileName = uuid.New().String()
		}
		fileName = fileName + data.Ext
		if (data.Type == "text") {
			fileName = fileName + ".txt"
		} else if (data.Type == "image") {
			fileName = fileName + ".png"
		}

		
		FtpHelper.UploadFile(*server, *username, *password, *port, *folder, *data, fileName)

		fmt.Println("File uploaded successfully")

		clipboard.Write(clipboard.FmtText, []byte(*webServer + "/" + fileName))
	}
	
	//WaitAndHandleShutdown()
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