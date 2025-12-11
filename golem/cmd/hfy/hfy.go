package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rle-lt/youtuber/golem/pkg/generation"
)

func main() {
	cwd, err := os.Getwd()
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

	assetsPath := filepath.Join(cwd, "assets/hfy/")
	templatePath := filepath.Join(assetsPath, "template.json")

	file, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var parsedJson generation.VideoOutlineGenerationTemplate

	err = json.Unmarshal(file, &parsedJson)
	if err != nil {
		log.Fatalf("failed to parse JSON file: %v", err)
	}

	promptCount := 5
	prompts, err := generator.GenerateVideoOutlines(parsedJson, uint(promptCount))
	if err != nil {
		log.Fatalf("Failed to generate title and rough outline: %v", err)

	}

	for _, p := range prompts {

		fmt.Printf(p.Title + "\n\n")
		fmt.Printf(p.Outline + "\n\n")

	}

	var input uint

	fmt.Print("Choose prompt to generate story\n")
	fmt.Scan(&input)

	if input == 0 {
		return
	}
	pick := input - 1

	fmt.Printf("Story picked: %s", prompts[pick].Title)
	pickedOutline := prompts[pick].Title

	generator.GenerateStory(pickedOutline)

}
