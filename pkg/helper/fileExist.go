package helper

import (
	"bufio"
	"fmt"
	"io"
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

func SaveFile(file *os.File) (string, error) {
	basePath := "uploads/video_"
	filePath := filepath.Join(basePath, file.Name())

	// Проверка существования файла и добавление цифры, если необходимо
	i := 1
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		filePath = filepath.Join(basePath, fmt.Sprintf("%s(%d)", file.Name(), i))
		i++
	}

	file , err := os.Create(filePath)
	if err!= nil {
          return "", err
     }
	defer file.Close()

	w := bufio.NewWriter(file)
	buf := make([]byte, 1024)
	for {
		n , err := file.Read(buf)
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
