package runner

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/rle-lt/youtuber/scripter/pkg/scripter"
)

type Tempalte struct {
	Idea                          string
	TargetAudienceCharacteristics []string
	ExamplePrompts                []string
	TitleStructureVariations      []string
}

func GenerateHFY(pathToTemplate string) {

	path := filepath.Join(pathToTemplate, "assets/hfy/template.json")

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error in reading json file: %v", err)
		return
	}

	var template Tempalte

	json.Unmarshal(data, &template)

	config := scripter.Config{
		APIKey: os.Getenv("OPENROUTER_API_KEY"),
		Models: scripter.Models{
			InitialOutline: "openrouter://amazon/nova-2-lite-v1:free",
			ChapterOutline: "openrouter://amazon/nova-2-lite-v1:free",
			ChapterWriter:  "openrouter://amazon/nova-2-lite-v1:free",
			Revision:       "openrouter://amazon/nova-2-lite-v1:free",
			Scrub:          "openrouter://amazon/nova-2-lite-v1:free",
			Prompt:         "openrouter://amazon/nova-2-lite-v1:free",
		},
		MaxChapterCount: 6,
		OutputWriter:    os.Stdout,
		StatusWriter:    os.Stderr,
	}

	generator, err := scripter.NewGenerator(config)
	if err != nil {
		log.Fatal(err)
	}
	generator.GeneratePrompts()

}
