package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rle-lt/youtuber/scripter/pkg/scripter"
)

func main() {
	// Configure the generator
	config := scripter.Config{
		APIKey: os.Getenv("OPENROUTER_API_KEY"),
		Models: scripter.Models{
			InitialOutline: "openrouter://amazon/nova-2-lite-v1:free",
			ChapterOutline: "openrouter://amazon/nova-2-lite-v1:free",
			ChapterWriter:  "openrouter://amazon/nova-2-lite-v1:free",
			Revision:       "openrouter://amazon/nova-2-lite-v1:free",
			Scrub:          "openrouter://amazon/nova-2-lite-v1:free",
		},
		MaxChapterCount: 6,
		OutputWriter:    os.Stdout,
		StatusWriter:    os.Stderr,
	}

	// Create generator
	generator, err := scripter.NewGenerator(config)
	if err != nil {
		log.Fatal(err)
	}

	// Generate story
	prompt := "Write a science fiction story about a lone astronaut..."
	chapters, err := generator.GenerateStory(prompt)
	if err != nil {
		log.Fatal(err)
	}

	// Output the story
	fmt.Println(strings.Join(chapters, "\n\n"))
}
