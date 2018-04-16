package main

import (
	"image"
	"image/jpeg"
	"io"
)

type iJpeg struct {
	ext string
}

func (j *iJpeg) getExt() string {
	return j.ext
}

func (j *iJpeg) decode(file io.Reader) (image.Image, error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (j *iJpeg) encode(out io.Writer, img image.Image) error {
	err := jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}
	return nil
}
