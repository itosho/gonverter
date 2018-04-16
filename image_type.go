package main

import (
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
)

type imageType interface {
	getExt() string
	decode(io.Reader) (image.Image, error)
	encode(io.Writer, image.Image) error
}

func getImageType(ext string) imageType {
	switch ext {
	case ".jpg", ".jpeg":
		return &iJpeg{ext}
	case ".png":
		return &iPng{ext}
	default:
		return nil
	}
}

func filePathWalk(directory string, fromType imageType, toType imageType) int {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		fromExt := fromType.getExt()

		if filepath.Ext(path) == fromExt {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			img, err := fromType.decode(file)
			if err != nil {
				return err
			}

			out, err := os.Create(path[:len(path)-len(fromExt)] + toType.getExt())
			if err != nil {
				return err
			}
			defer out.Close()

			toType.encode(out, img)

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
