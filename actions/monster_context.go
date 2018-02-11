package actions

import (
	"github.com/gobuffalo/buffalo"
	"time"
	"io"
	"path"

	"github.com/pkg/errors"
	"net/http"
	"os"
	"strconv"
)


type MonsterContext struct {
	buffalo.Context
}

func (m *MonsterContext) RenderFile(file string) error {
	start := time.Now()
	defer func() {
		m.LogField("render", time.Since(start))
	}()

	p := path.Join("data", file)
	_, fileName := path.Split(file)

	//Check if file exists and open
	oFile, err := os.Open(p)
	defer oFile.Close() //Close after function return

	if err != nil {
		//File not found, send 404
		return buffalo.HTTPError{Status: 404, Cause: errors.New("file not found")}
	}

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	oFile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := oFile.Stat()                        //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	m.Response().Header().Set("Content-Type", FileContentType)
	m.Response().Header().Set("Content-Disposition", "attachment; filename="+fileName)
	m.Response().Header().Set("Content-Length", FileSize)
	m.Response().WriteHeader(200)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	oFile.Seek(0, 0)

	_, err = io.Copy(m.Response(), oFile) //'Copy' the file to the client
	if err != nil {
		return buffalo.HTTPError{Status: 500, Cause: errors.WithStack(err)}
	}

	return nil
}