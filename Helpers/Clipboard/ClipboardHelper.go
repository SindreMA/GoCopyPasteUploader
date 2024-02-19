package ClipboardHelper

import (
	"main/Types"
	"strings"

	"golang.design/x/clipboard"
)

func hasFileExtension(s string) bool {
	lastDotIndex := strings.LastIndex(s, ".")
	// Ensure there's a dot, it's not the last character, and the extension is not more than 10 characters long
	return lastDotIndex > -1 && lastDotIndex < len(s)-1 && len(s)-lastDotIndex-1 <= 20
}

func HandleFileClipboard(textClipboard string) *Types.IClipboard {
	fileClipboard := GetFileClipboard()

	if fileClipboard != "" {
		return &Types.IClipboard {
			Type: Types.File,
			Data: string(fileClipboard),
			Ext: GetFileExtension(textClipboard),
		}
	}

	return nil
}

func GetClipboardData() *Types.IClipboard {

	textClipboard := clipboard.Read(clipboard.FmtText)
	
	if textClipboard != nil && len(textClipboard) > 0 {
		textClipboardString := string(textClipboard)

		if hasFileExtension(textClipboardString) {
			fileClipboard := HandleFileClipboard(textClipboardString)

			if fileClipboard != nil {
				return fileClipboard
			}
		}
		if textClipboardString != "ClipboardFile"{
			return &Types.IClipboard {
				Type: Types.Text,
				Data: string(textClipboard),
			}
		}
	}

	imageClipboard := GetImageClipboard()

	if imageClipboard != "" {
		return &Types.IClipboard {
			Type: Types.Image,
			Data: string(imageClipboard),
		}
	}

	fileClipboard := HandleFileClipboard(string(textClipboard))

	if fileClipboard != nil {
		return fileClipboard
	}
	
	return nil
}

func GetFileExtension(s string) string {
	lastDotIndex := strings.LastIndex(s, ".")
	return s[lastDotIndex:]
}