package main

import (
	"image"
	"io"
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
