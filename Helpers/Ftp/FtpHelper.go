package FtpHelper

import (
	"encoding/base64"
	"main/Types"
	"os"
	"path/filepath"

	ftp "github.com/martinr92/goftp"
)

func UploadFile(server string, username string, password string, port int, folder string, data Types.IClipboard, fileName string) {
	ftpClient, err := ftp.NewFtp( server + ":" + string(port))
	if err != nil {
		panic(err)
	}
	ftpClient.Login(username, password)
	newFolder := filepath.Join(folder)

	// Temporary file path store
	tempFilePath := "/tmp/" + fileName

	// Create a write base64 data to file
	file, err := os.Create(tempFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if (data.Type == "text") {
		file.Write([]byte(data.Data))
	} else {
		file.Write(Base64Decode(data.Data))
	}
	
	ftpClient.Upload(tempFilePath, newFolder + "/" + fileName)

	// Remove the temporary file
	os.Remove(tempFilePath)

	ftpClient.Close()
}

func Base64Decode(data string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return decoded
}