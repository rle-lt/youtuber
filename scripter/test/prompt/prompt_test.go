package test

import (
	"os"
	"testing"

	"github.com/rle-lt/youtuber/scripter/pkg/scripter"
)

func TestPromptGeneration(t *testing.T) {
	config := scripter.Config{
		APIKey: os.Getenv("OPENROUTER_API_KEY"),
		Models: scripter.Models{
			Prompt: "openrouter://amazon/nova-2-lite-v1:free",
		},
	}

	generator, err := scripter.NewGenerator(config)
	if err != nil {
		t.Errorf(`Error in creating a new Generator with err\n %v`, err)

	}

	prompt := "Stories about mediaval times"
	_, err = generator.GeneratePrompts(prompt, []string{""}, []string{""}, []string{""}, 2)
	if err != nil {
		t.Errorf(`Error in generating test prompts with err\n %v`, err)
	}
}
