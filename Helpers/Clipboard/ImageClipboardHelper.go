package ClipboardHelper

import (
	"encoding/base64"

	"golang.design/x/clipboard"
)

func GetImageClipboard() string {
	imageClipboard := clipboard.Read(clipboard.FmtImage)

	if imageClipboard == nil || len(imageClipboard) == 0 {
		return ""
	}
		
	return ImageFromByteArrayToBase64(imageClipboard)
}

func ImageFromByteArrayToBase64(imageByteArray []byte) string {
	base64String := base64.StdEncoding.EncodeToString(imageByteArray)
	return base64String
}