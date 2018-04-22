package jpeg

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/itosho/gonverter/cli"
)

type Jpeg struct{}

const (
	Quality = 100
)

func init() {
	cli.Register(".jpeg", Jpeg{})
	cli.Register(".jpg", Jpeg{})
}

func (j Jpeg) Decode(file io.Reader) (image.Image, error) {
	return jpeg.Decode(file)
}

func (j Jpeg) Encode(out io.Writer, img image.Image) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: Quality})
}