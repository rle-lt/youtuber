package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rle-lt/youtuber/julien/pkg/generation"
)

func main() {
	pathToFile, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config := generation.Config{
		APIKey: os.Getenv("OPENROUTER_API_KEY"),
		Models: generation.Models{
			Prompt: "openrouter://amazon/nova-2-lite-v1:free",
		},
		MaxSceneCount: 6,
		StatusWriter:  os.Stderr,
	}

	generator, err := generation.NewGenerator(config)
	if err != nil {
		log.Fatal(err)
	}

	pathToFile = filepath.Join(pathToFile, "assets/hfy/template.json")

	file, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var parsedJson generation.PromptGenerationTemplate

	err = json.Unmarshal(file, &parsedJson)
	if err != nil {
		log.Fatalf("failed to parse JSON file: %v", err)
	}

	prompts, err := generator.GeneratePrompts(parsedJson, 5)
	if err != nil {
		log.Fatalf("Failed to generate title and rough outline: %v", err)

	}

	for _, p := range prompts {
		fmt.Printf(p.Title + "\n")
		fmt.Printf(p.Outline + "\n")
	}
}
