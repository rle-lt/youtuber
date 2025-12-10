package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rle-lt/youtuber/runner/pkg/runner"
)

func main() {
	pathToFile, err := os.Getwd()
	if err != nil {
		log.Fatal(err)

	}

	pathToFile = filepath.Join(pathToFile, "assets/hfy")

	runner.GenerateHFY(pathToFile)

}
