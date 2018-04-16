package main

import (
	"flag"
	"log"
	"os"
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
