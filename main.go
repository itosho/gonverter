package main

import (
	"flag"
	"fmt"
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
	var toExt = flag.String("t", ".png", "to extension")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Invalid Args. Please specify only one direcoty.")
	}

	args := flag.Args()
	directory := args[0]

	fromType := getImageType(*fromExt)
	if fromType == nil {
		log.Fatalln("Invalid extenstion type.")
	}

	toType := getImageType(*toExt)
	if toType == nil {
		log.Fatalln("Invalid extenstion type.")
	}

	code := filePathWalk(directory, fromType, toType)
	os.Exit(code)
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
