package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/itosho/gonverter/con"
)

const (
	ExitSuccess int = iota
	ExitError
	ExitFileError
)

var images = map[string]con.DecodeEncoder{}

func Register(ext string, image con.DecodeEncoder) {
	images[ext] = image
}

func Run() {
	var fromExt = flag.String("f", ".jpg", "from extension")
	var toExt = flag.String("t", ".png", "to extension")
	flag.Usage = usage

	if flag.NArg() != 1 {
		log.Fatal("Invalid Args. Please specify only one direcoty.")
	}

	args := flag.Args()
	directory := args[0]

	code, err := convert(directory, fromExt, toExt)
	os.Exit(code)
}

func generateConvertFile(path string, fromExt string, toExt string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fromImage, ok := images[fromExt]
	if ok {
		img, err := fromImage.Decode(file)
		if err != nil {
			return err
		}
	}

	convertFilePath := getConvertFilePath(path, fromExt, toExt)
	out, err := os.Create(convertFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	toImage, ok := images[toExt]
	if ok {
		img, err := toImage.Encode(out, img)
		if err != nil {
			return err
		}
	}
	return nil
}

func convert(directory string, fromExt string, toExt string) bool {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) == fromExt {
			err := generateConvertFile(path, fromExt, toExt)
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
		fmt.Fprintln(os.Stderr, "Convert Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	return true
}

func getConvertFilePath(path string, fromExt string, toExt string) string {
	return path[:len(path)-len(fromExt)] + toExt
}

func usage() {
	fmt.Println("usage: gonverter [-f from extension] [-t to extension] [directory]")
	flag.PrintDefaults()
	os.Exit(ExitSuccess)
}
