package main

import (
	"image"
	"image/png"
	"io"
)

type iPng struct {
	ext string
}

func (p *iPng) getExt() string {
	return p.ext
}

func (p *iPng) decode(file io.Reader) (image.Image, error) {
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (p *iPng) encode(out io.Writer, img image.Image) error {
	err := png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}
