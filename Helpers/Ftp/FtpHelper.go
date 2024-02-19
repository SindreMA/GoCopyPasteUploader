package FtpHelper

import (
	"encoding/base64"
	"main/Types"
	"os"
	"path/filepath"
	"strconv"

	ftp "github.com/martinr92/goftp"
)

func UploadFile(server string, username string, password string, port int, folder string, data Types.IClipboard, fileName string) {
	ftpClient, err := ftp.NewFtp( server + ":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	ftpClient.Login(username, password)
	newFolder := filepath.Join(folder)

	dir, err := os.Getwd()
    if err != nil {
        panic(err) // or handle the error in a way that makes sense for your application
    }

	// Temporary file path store
	tempFilePath := filepath.Join(dir, "tmp", fileName)

	// create folder if not exists
	if _, err := os.Stat(filepath.Join(dir, "tmp")); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dir, "tmp"), 0755)
	}

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

	
	
	ftpClient.Upload(tempFilePath, filepath.Join(newFolder, fileName))

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