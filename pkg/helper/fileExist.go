package helper

import (
	"bufio"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

func SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {


	basePath := "uploads/video_"
	filePath := filepath.Join(basePath, header.Filename)

	i := 1
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		filePath = filepath.Join(basePath, fmt.Sprintf("%s(%d)", header.Filename, i))
		i++
	}

	fileNew , err := os.Create(filePath)
	if err!= nil {
          return "", err
     }
	defer fileNew.Close()

	w := bufio.NewWriter(fileNew)
	buf := make([]byte, 1024)
	for {
		n , err := fileNew.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
               break
          }

		if _, err := w.Write(buf[:n]); err!= nil {
			return "", err
		}
	}

	if err := w.Flush(); err!= nil {
		panic(err)
	}

	return filePath, nil
}
