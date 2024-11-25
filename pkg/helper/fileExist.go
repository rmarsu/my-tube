package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/VandiKond/vanerrors"
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
	basePath := "uploads/"
	filePath := filepath.Join(basePath, "video_"+header.Filename)

	i := 1
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		ext := filepath.Ext(header.Filename)
		name := header.Filename[:len(header.Filename)-len(ext)]
		filePath = filepath.Join(basePath, fmt.Sprintf("video_%s(%d)%s", name, i, ext))
		i++
	}

	fileNew, err := os.Create(filePath)
	if err != nil {
		err = vanerrors.NewWrap("error to create a file", err, vanerrors.EmptyHandler)
		return "", err
	}
	defer fileNew.Close()

	// Implementing saving
	if _, err := io.Copy(fileNew, file); err != nil {
		err = vanerrors.NewWrap("error to copy data", err, vanerrors.EmptyHandler)
		return "", err
	}

	return filePath, nil
}
