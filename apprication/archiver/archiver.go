package archiver

import (
	"github.com/jlaffaye/ftp"
	"io"
	"os"
	"log"
)

// Archive does stuff...
func Archive(filename string, response *ftp.Response) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	log.Println("Writing file:", filename)
	io.Copy(file, response)
}
