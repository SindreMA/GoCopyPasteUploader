package ClipboardHelper

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetFileClipboard() string {
	
	filePath := GetFileClipboardData()
	
	if filePath == "" {
		return ""
	}
	
	// if file doesn't exist, return empty string
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return ""
	}

	if fileInfo, err := os.Stat(filePath); err == nil && fileInfo.IsDir() {
		return ""
	}
	
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
		return ""
	}

	// Encode the file content to Base64
	encodedString := base64.StdEncoding.EncodeToString(data)

	return encodedString
}

func isMacOS() bool {
    return runtime.GOOS == "darwin"
}

func GetFileClipboardData() string {
	// If is Mac
	if isMacOS() {
		return GetFileClipboardMac()
	}
	// If is Windows
	return GetFileClipboardWindows()
}

func GetFileClipboardWindows() string {
	// Not implemented yet
	return "";
}

func GetFileClipboardMac() string {
	cmd := exec.Command("./External/MacFileClipboard")
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error executing command:", err)
        return ""
    }

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		return line
	}

	return "";
}