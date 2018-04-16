package convert

import (
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
)

const (
	ExitSuccess int = iota
	ExitError
	ExitFileError
)

type ConvertType interface {
	GetExt() string
	Decode(io.Reader) (image.Image, error)
	Encode(io.Writer, image.Image) error
}

func GetImageType(ext string) ConvertType {
	switch ext {
	case ".jpg", ".jpeg":
		return &cjpg{ext}
	case ".png":
		return &cpng{ext}
	default:
		return nil
	}
}

func Convert(directory string, fromType ConvertType, toType ConvertType) int {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) == fromType.GetExt() {
			err := generateConvertFile(path, fromType, toType)
			if err != nil {
				return err
			}

			// remove originalFile
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitError
	}

	return ExitSuccess
}

func generateConvertFile(path string, fromType ConvertType, toType ConvertType) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := fromType.Decode(file)
	if err != nil {
		return err
	}

	convertFilePath := getConvertFilePath(path, fromType.GetExt(), toType.GetExt())
	out, err := os.Create(convertFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	toType.Encode(out, img)
	return nil
}

func getConvertFilePath(path string, fromExt string, toExt string) string {
	return path[:len(path)-len(fromExt)] + toExt
}
