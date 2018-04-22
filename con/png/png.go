package png

import (
	"image"
	"image/png"
	"io"

	"github.com/itosho/gonverter/cli"
)

type Png struct{}

func init() {
	cli.Register(".png", Png{})
}

func (p Png) Decode(file io.Reader) (image.Image, error) {
	return png.Decode(file)
}

func (p Png) Encode(out io.Writer, img image.Image) error {
	return png.Encode(out, img)
}
