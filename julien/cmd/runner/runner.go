package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rle-lt/youtuber/julien/pkg/hfy"
)

func main() {
	pathToFile, err := os.Getwd()
	if err != nil {
		log.Fatal(err)

	}

	pathToFile = filepath.Join(pathToFile, "assets/hfy")

	hfy.GenerateHFY(pathToFile)

}
