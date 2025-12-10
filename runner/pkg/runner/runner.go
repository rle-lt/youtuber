package runner

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rle-lt/youtuber/scripter/pkg/scripter"
)

func GenerateHFY(pathToTemplate string) {

	path := filepath.Join(pathToTemplate, "template.json")

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error in reading json file: %v", err)
		return
	}

	var template scripter.PromptGenerationTemplate

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
		MaxSceneCount: 6,
		StatusWriter:  os.Stderr,
	}

	generator, err := scripter.NewGenerator(config)
	if err != nil {
		log.Fatal(err)
	}
	prompts, err := generator.GeneratePrompts(template, 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The prompts:\n %s", strings.Join(prompts, "\n"))

}
