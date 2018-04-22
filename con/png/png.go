package png

import (
	"image"
	"image/png"
	"io"
)

type Png struct{}

func (p Png) Decode(file io.Reader) (image.Image, error) {
	return png.Decode(file)
}

func (p Png) Encode(out io.Writer, img image.Image) error {
	return png.Encode(out, img)
}
