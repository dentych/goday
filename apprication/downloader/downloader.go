package downloader

import (
	"fmt"
	"github.com/jlaffaye/ftp"
)

func Hello() {
	fmt.Println("Hello!")
}

// Will download a file from an FTP server, specified by addr. Addr should contain a hostname and a port number.
// Filepath path and filename of the file to download.
// The method returns a Response which can be read for the contents of the downloaded file. The Response MUST BE closed
// after use, otherwise FTP errors might happen.
func DownloadFileAnonymously(addr, filePath string) *ftp.Response {
	return DownloadFile(addr, "anonymous", "anonymous", filePath)
}

func DownloadFile(addr, username, password, filePath string) *ftp.Response {
	con, err := ftp.Connect(addr)
	if err != nil {
		panic(err.Error())
	}
	con.Login(username, password)
	defer con.Logout()

	response, err := con.Retr(filePath)
	if err != nil {
		panic(err.Error())
	}

	return response
}
