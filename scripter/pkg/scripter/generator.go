package scripter

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	prompts "github.com/rle-lt/youtuber/scripter/pkg/prompt"
)

type Models struct {
	InitialOutline string
	ChapterOutline string
	ChapterWriter  string
	Revision       string
	Scrub          string
	Prompt         string
}

type Config struct {
	APIKey        string
	Models        Models
	MaxSceneCount uint
	StatusWriter  io.Writer
}

type Generator struct {
	config  Config
	clients map[string]*Client
}
type PromptGenerationTemplate struct {
	Idea                          string
	TargetAudienceCharacteristics []string
	ExamplePrompts                []string
	TitleStructureVariations      []string
}

func NewGenerator(config Config) (*Generator, error) {
	g := &Generator{
		config:  config,
		clients: make(map[string]*Client),
	}

	models := []string{
		config.Models.InitialOutline,
		config.Models.ChapterOutline,
		config.Models.ChapterWriter,
		config.Models.Revision,
		config.Models.Scrub,
		config.Models.Prompt,
	}

	for _, modelStr := range models {
		modelInfo, err := GetModelAndProvider(modelStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse model %s: %w", modelStr, err)
		}

		if _, exists := g.clients[modelInfo.Model]; !exists {
			if modelInfo.Provider == "openrouter" {
				client := NewClient(config.APIKey)
				g.clients[modelInfo.Model] = client
			}
		}
	}

	return g, nil
}

func (g *Generator) GeneratePrompts(idea string, titleStructure []string, usedPrompts []string, targetAudienceCharacteristics []string, resultExamples []string, promptCount uint) ([]string, error) {
	examplesStr := strings.Join(resultExamples, "\n")
	usedPromptsStr := ""
	audienceStr := ""
	titleStructureStr := ""
	for _, char := range targetAudienceCharacteristics {
		audienceStr += fmt.Sprintf("- %s\n", char)
	}
	for _, used := range usedPrompts {
		usedPromptsStr += fmt.Sprintf("- %s\n", used)
	}
	for _, title := range titleStructure {
		usedPromptsStr += fmt.Sprintf("- %s\n", title)
	}

	prompt := fmt.Sprintf(prompts.PROMPT_GENERATION_PROMPT,
		promptCount,
		idea,
		titleStructureStr,
		examplesStr,
		usedPromptsStr,
		audienceStr)

	modelInfo, err := GetModelAndProvider(g.config.Models.Prompt)

	if err != nil {
		return []string{}, err
	}

	messages := RequestMessages{
		{
			Role:    "system",
			Content: prompts.PROMPT_GENERATION_INTRO,
		},
		BuildMessage(prompt)}

	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return []string{}, err
	}

	generatedPrompts := messages.GetLastMessage()

	prompt = fmt.Sprintf(prompts.PROMPT_GENERATION_TO_JSON_PROMPT, generatedPrompts)

	messages = RequestMessages{BuildMessage(prompt)}

	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return []string{}, err
	}

	var parsedPrompts struct {
		Prompts []string `json:"prompts"`
	}
	err = json.Unmarshal([]byte(messages.GetLastMessage()), &parsedPrompts)
	if err != nil {
		return []string{}, err
	}
	return parsedPrompts.Prompts, nil
}

func (g *Generator) GenerateStory(prompt string) ([]string, error) {
	g.logStatus("Generating story outline...")
	storyOutline, storyElements, sceneOutline, importantInfo, err := g.GenerateInitOutline(prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate initial outline: %w", err)
	}
	g.logStatus("Finished story outline")

	g.logStatus("Reviewing story outline...")
	storyOutline, err = g.ReviewOutline(storyOutline)
	if err != nil {
		return nil, fmt.Errorf("failed to review outline: %w", err)
	}
	g.logStatus("Reviewed story outline")

	sceneCount, err := g.CountScenes(sceneOutline)
	if err != nil {
		return nil, fmt.Errorf("failed to count scenes: %w", err)
	}
	g.logStatus(fmt.Sprintf("Scene count is %d", sceneCount))

	g.logStatus("Generating scene outlines...")
	sceneOutlines, err := g.GenerateSceneOutlines(storyOutline, sceneCount)
	if err != nil {
		return nil, fmt.Errorf("failed to generate scene outlines: %w", err)
	}
	g.logStatus("Finished scene outlines")

	g.logStatus("Generating scenes...")
	scenes, err := g.GenerateScenes(sceneOutlines, sceneCount, storyOutline, importantInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to generate scenes: %w", err)
	}
	g.logStatus("Finished scenes")

	g.logStatus("Reviewing scenes...")
	scenes, err = g.ReviewScenes(scenes)
	if err != nil {
		return nil, fmt.Errorf("failed to review scenes: %w", err)
	}
	g.logStatus("Reviewed scenes")

	g.logStatus("Scrubbing scenes...")
	scenes, err = g.ScrubScenes(scenes)
	if err != nil {
		return nil, fmt.Errorf("failed to scrub scenes: %w", err)
	}
	g.logStatus("Scrubbed scenes")

	_ = storyElements

	return scenes, nil
}

func (g *Generator) GenerateInitOutline(outlinePrompt string) (string, string, string, string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.InitialOutline)
	if err != nil {
		return "", "", "", "", err
	}

	prompt := fmt.Sprintf(prompts.GET_IMPORTANT_BASE_PROMPT_INFO, outlinePrompt)
	messages := RequestMessages{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", "", "", "", err
	}
	importantInfo := messages.GetLastMessage()

	prompt = fmt.Sprintf(prompts.STORY_ELEMENTS_PROMPT, outlinePrompt)
	messages = RequestMessages{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", "", "", "", err
	}
	storyElements := messages.GetLastMessage()

	prompt = fmt.Sprintf(prompts.INITIAL_OUTLINE_PROMPT, outlinePrompt, storyElements, g.config.MaxSceneCount)
	messages = RequestMessages{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", "", "", "", err
	}
	initialOutline := messages.GetLastMessage()

	fullOutline := initialOutline + storyElements + importantInfo
	return fullOutline, storyElements, initialOutline, importantInfo, nil
}

func (g *Generator) CountScenes(sceneOutline string) (uint, error) {
	prompt := fmt.Sprintf(prompts.SCENE_COUNT_PROMPT, sceneOutline)
	modelInfo, err := GetModelAndProvider(g.config.Models.ChapterOutline)
	if err != nil {
		return 0, err
	}

	messages := RequestMessages{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return 0, err
	}

	output := messages.GetLastMessage()
	var data struct {
		TotalChapters int `json:"totalChapters"`
	}

	err = json.Unmarshal([]byte(output), &data)
	if err != nil {
		return 0, fmt.Errorf("unable to unmarshal scene count: %w", err)
	}

	return uint(data.TotalChapters), nil
}

func (g *Generator) GenerateSceneOutlines(outline string, outlineCount uint) ([]string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.ChapterOutline)
	if err != nil {
		return nil, err
	}

	prompt := fmt.Sprintf("Please help me expand upon the following outline, scene by scene.\n%s", outline)
	messages := RequestMessages{BuildMessage(prompt)}
	sceneOutlines := make([]string, 0, outlineCount)

	for i := range outlineCount {
		sceneNum := i + 1
		prompt := fmt.Sprintf(prompts.SCENE_OUTLINE_PROMPT, sceneNum, outline, sceneNum)
		messages = append(messages, BuildMessage(prompt))

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate scene %d outline: %w", sceneNum, err)
		}

		sceneOutlines = append(sceneOutlines, messages.GetLastMessage())
	}

	return sceneOutlines, nil
}

func (g *Generator) GenerateScenes(sceneOutlines []string, sceneCount uint, storyOutline string, importantInfo string) ([]string, error) {
	scenes := make([]string, 0, sceneCount)
	modelInfo, err := GetModelAndProvider(g.config.Models.ChapterWriter)
	if err != nil {
		return nil, err
	}

	for i := range sceneCount {
		sceneNum := i + 1
		currentSceneOutline := sceneOutlines[i]

		previousScenes := ""
		for j := range i {
			previousScenes += scenes[j] + "\n"
		}

		lastSceneSummary := ""
		if i > 0 {
			prompt := fmt.Sprintf(prompts.SCENE_SUMMARY_PROMPT, sceneNum, storyOutline, scenes[i-1])
			messages := RequestMessages{
				{Role: "system", Content: prompts.SCENE_SUMMARY_INTRO},
				BuildMessage(prompt),
			}

			messages, err = g.generateText(messages, modelInfo.Model)
			if err != nil {
				return nil, fmt.Errorf("failed to generate scene %d summary: %w", sceneNum, err)
			}
			lastSceneSummary = messages.GetLastMessage()
		}

		previousScenesContext := ""
		if previousScenes != "" {
			previousScenesContext = fmt.Sprintf("Previous scenes context:\n%s", previousScenes)
		}

		prompt := fmt.Sprintf(prompts.SCENE_GENERATION_STAGE1,
			previousScenesContext, importantInfo, sceneNum, sceneCount,
			currentSceneOutline, lastSceneSummary, sceneNum)

		messages := RequestMessages{
			{Role: "system", Content: fmt.Sprintf(prompts.SCENE_GENERATION_INTRO, storyOutline)},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate scene %d stage 1: %w", sceneNum, err)
		}
		firstSceneStage := messages.GetLastMessage()

		prompt = fmt.Sprintf(prompts.SCENE_GENERATION_STAGE2,
			previousScenesContext, importantInfo, sceneNum, sceneCount,
			currentSceneOutline, lastSceneSummary, firstSceneStage, sceneNum)

		messages = RequestMessages{
			{Role: "system", Content: fmt.Sprintf(prompts.SCENE_GENERATION_INTRO, storyOutline)},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate scene %d stage 2: %w", sceneNum, err)
		}
		secondSceneStage := messages.GetLastMessage()

		prompt = fmt.Sprintf(prompts.SCENE_GENERATION_STAGE3,
			previousScenesContext, importantInfo, sceneNum, sceneCount,
			lastSceneSummary, secondSceneStage, sceneNum)

		messages = RequestMessages{
			{Role: "system", Content: fmt.Sprintf(prompts.SCENE_GENERATION_INTRO, storyOutline)},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate scene %d stage 3: %w", sceneNum, err)
		}
		thirdSceneStage := messages.GetLastMessage()

		scenes = append(scenes, thirdSceneStage)
	}

	return scenes, nil
}

func (g *Generator) ReviewOutline(storyOutline string) (string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.Revision)
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf(prompts.CRITIC_OUTLINE_PROMPT, storyOutline)
	messages := RequestMessages{
		{Role: "system", Content: prompts.CRITIC_OUTLINE_INTRO},
		BuildMessage(prompt),
	}

	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", err
	}
	feedback := messages.GetLastMessage()

	prompt = fmt.Sprintf(prompts.OUTLINE_REVISION_PROMPT, storyOutline, feedback)
	messages = RequestMessages{BuildMessage(prompt)}

	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", err
	}

	return messages.GetLastMessage(), nil
}

func (g *Generator) ReviewScenes(scenes []string) ([]string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.Revision)
	if err != nil {
		return nil, err
	}

	reviewedScenes := make([]string, 0, len(scenes))
	for i, scene := range scenes {
		prompt := fmt.Sprintf(prompts.CRITIC_SCENE_PROMPT, scene)
		messages := RequestMessages{
			{Role: "system", Content: prompts.CRITIC_SCENE_INTRO},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to get feedback for scene %d: %w", i+1, err)
		}
		feedback := messages.GetLastMessage()

		prompt = fmt.Sprintf(prompts.SCENE_REVISION_PROMPT, scene, feedback)
		messages = RequestMessages{BuildMessage(prompt)}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to revise scene %d: %w", i+1, err)
		}

		reviewedScenes = append(reviewedScenes, messages.GetLastMessage())
	}

	return reviewedScenes, nil
}

func (g *Generator) ScrubScenes(scenes []string) ([]string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.Scrub)
	if err != nil {
		return nil, err
	}

	scrubbedScenes := make([]string, 0, len(scenes))
	for i, scene := range scenes {
		prompt := fmt.Sprintf(prompts.SCRUB_PROMPT, scene)
		messages := RequestMessages{BuildMessage(prompt)}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to scrub scene %d: %w", i+1, err)
		}

		scrubbedScenes = append(scrubbedScenes, messages.GetLastMessage())
	}

	return scrubbedScenes, nil
}

func (g *Generator) generateText(messages RequestMessages, model string) (RequestMessages, error) {
	client, ok := g.clients[model]
	if !ok {
		return nil, fmt.Errorf("no client configured for model: %s", model)
	}

	return client.GenerateText(messages, model)
}

func (g *Generator) logStatus(message string) {
	if g.config.StatusWriter != nil {
		fmt.Fprintf(g.config.StatusWriter, "%s\n", message)
	}
}
