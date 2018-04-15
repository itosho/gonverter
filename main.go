package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

const (
	ExitSuccess int = iota
	ExitError
	ExitFileError
)

func main() {
	var fromExt = flag.String("f", ".jpg", "from extension")
	// var toExt = flag.String("t", "png", "to extension")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Invalid Args. Please specify only one direcoty.")
	}

	args := flag.Args()
	directory := args[0]

	code := filePathWalk(directory, *fromExt)
	os.Exit(code)
}

func filePathWalk(directory string, fromExt string) int {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == fromExt {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			img, err := jpeg.Decode(file)
			if err != nil {
				fmt.Println(path)
				return err
			}

			out, err := os.Create(path[:len(path)-len(fromExt)] + ".png")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}
			defer out.Close()

			png.Encode(out, img)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitError
	}

	return ExitSuccess
}
