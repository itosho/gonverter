package cli

import "github.com/itosho/gonverter/con"

var images = map[string]con.DecodeEncoder{}

func Register(ext string, image con.DecodeEncoder) {
	images[ext] = image
}
