package test

import (
	"os"
	"testing"

	"github.com/rle-lt/youtuber/scripter/pkg/scripter"
)

func TestStoryGeneration(t *testing.T) {
	config := scripter.Config{
		APIKey: os.Getenv("OPENROUTER_API_KEY"),
		Models: scripter.Models{
			InitialOutline: "openrouter://amazon/nova-2-lite-v1:free",
			ChapterOutline: "openrouter://amazon/nova-2-lite-v1:free",
			ChapterWriter:  "openrouter://amazon/nova-2-lite-v1:free",
			Revision:       "openrouter://amazon/nova-2-lite-v1:free",
			Scrub:          "openrouter://amazon/nova-2-lite-v1:free",
		},
		MaxChapterCount: 1,
		StatusWriter:    os.Stderr,
	}

	generator, err := scripter.NewGenerator(config)
	if err != nil {
		t.Errorf(`Error in creating a new Generator with err\n %v`, err)

	}

	prompt := "Write a science fiction story about a lone astronaut"
	_, err = generator.GenerateStory(prompt)
	if err != nil {
		t.Errorf(`Error in generating test story with err\n %v`, err)
	}
}
