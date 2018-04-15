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
	var fromExt = flag.String("f", ".go", "from extension")
	// var toExt = flag.String("t", "png", "to extension")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Invalid Args. Please specify only one direcoty.")
	}

	args := flag.Args()
	directory := args[0]

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == *fromExt {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(ExitSuccess)
}
