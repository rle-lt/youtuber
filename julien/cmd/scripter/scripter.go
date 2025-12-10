package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	defaults "github.com/rle-lt/youtuber/scripter/cmd/default"
	"github.com/rle-lt/youtuber/scripter/pkg/scripter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v\n", err)
	}

	promptPath := flag.String("prompt", "", "Path to prompt file. -prompt /path/to/file")
	flag.Parse()

	if *promptPath == "" {
		log.Fatal("Please specify a prompt file with -prompt /path/to/file")
	}

	prompt, err := os.ReadFile(*promptPath)
	if err != nil {
		log.Fatalf("Error reading prompt file: %v", err)
	}

	if len(prompt) == 0 {
		log.Fatal("Prompt file is empty. Please provide story context.")
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENROUTER_API_KEY not set in environment")
	}

	config := scripter.Config{
		APIKey: apiKey,
		Models: scripter.Models{
			InitialOutline: defaults.INITIAL_OUTLINE_WRITER_MODEL,
			ChapterOutline: defaults.CHAPTER_OUTLINE_WRITER_MODEL,
			ChapterWriter:  defaults.CHAPTER_WRITER_MODEL,
			Revision:       defaults.REVISION_MODEL,
			Scrub:          defaults.SCRUB_MODEL,
		},
		MaxChapterCount: defaults.MAX_CHAPTER_COUNT,
		StatusWriter:    os.Stderr,
	}

	generator, err := scripter.NewGenerator(config)
	if err != nil {
		log.Fatalf("Failed to create generator: %v", err)
	}

	chapters, err := generator.GenerateStory(string(prompt))
	if err != nil {
		log.Fatalf("Failed to generate story: %v", err)
	}

	// Output the final story
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(chapters, "\n"))
}
