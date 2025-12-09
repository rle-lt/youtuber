package scripter

import (
	"encoding/json"
	"fmt"
	"io"

	prompts "github.com/rle-lt/youtuber/scripter/pkg/prompt"
)

type Models struct {
	InitialOutline string
	ChapterOutline string
	ChapterWriter  string
	Revision       string
	Scrub          string
}

type Config struct {
	APIKey          string
	Models          Models
	MaxChapterCount uint
	StatusWriter    io.Writer
}

type Generator struct {
	config  Config
	clients map[string]*Client
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

func (g *Generator) GenerateStory(prompt string) ([]string, error) {
	// Generate initial outline
	g.logStatus("Generating story outline...")
	storyOutline, storyElements, chapterOutline, importantInfo, err := g.GenerateInitOutline(prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate initial outline: %w", err)
	}
	g.logStatus("Finished story outline")

	// Review outline
	g.logStatus("Reviewing story outline...")
	storyOutline, err = g.ReviewOutline(storyOutline)
	if err != nil {
		return nil, fmt.Errorf("failed to review outline: %w", err)
	}
	g.logStatus("Reviewed story outline")

	// Count chapters
	chapterCount, err := g.CountChapters(chapterOutline)
	if err != nil {
		return nil, fmt.Errorf("failed to count chapters: %w", err)
	}
	g.logStatus(fmt.Sprintf("Chapter count is %d", chapterCount))

	// Generate chapter outlines
	g.logStatus("Generating chapter outlines...")
	chapterOutlines, err := g.GenerateChapterOutlines(storyOutline, chapterCount)
	if err != nil {
		return nil, fmt.Errorf("failed to generate chapter outlines: %w", err)
	}
	g.logStatus("Finished chapter outlines")

	// Generate chapters
	g.logStatus("Generating chapters...")
	chapters, err := g.GenerateChapters(chapterOutlines, chapterCount, storyOutline, importantInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to generate chapters: %w", err)
	}
	g.logStatus("Finished chapters")

	// Review chapters
	g.logStatus("Reviewing chapters...")
	chapters, err = g.ReviewChapters(chapters)
	if err != nil {
		return nil, fmt.Errorf("failed to review chapters: %w", err)
	}
	g.logStatus("Reviewed chapters")

	// Scrub chapters
	g.logStatus("Scrubbing chapters...")
	chapters, err = g.ScrubChapters(chapters)
	if err != nil {
		return nil, fmt.Errorf("failed to scrub chapters: %w", err)
	}
	g.logStatus("Scrubbed chapters")

	_ = storyElements

	return chapters, nil
}

func (g *Generator) GenerateInitOutline(outlinePrompt string) (string, string, string, string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.InitialOutline)
	if err != nil {
		return "", "", "", "", err
	}

	prompt := fmt.Sprintf(prompts.GET_IMPORTANT_BASE_PROMPT_INFO, outlinePrompt)
	messages := []RequestMessage{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", "", "", "", err
	}
	importantInfo := GetLastMessage(messages)

	prompt = fmt.Sprintf(prompts.STORY_ELEMENTS_PROMPT, outlinePrompt)
	messages = []RequestMessage{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", "", "", "", err
	}
	storyElements := GetLastMessage(messages)

	prompt = fmt.Sprintf(prompts.INITIAL_OUTLINE_PROMPT, outlinePrompt, storyElements, g.config.MaxChapterCount)
	messages = []RequestMessage{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", "", "", "", err
	}
	initialOutline := GetLastMessage(messages)

	fullOutline := initialOutline + storyElements + importantInfo
	return fullOutline, storyElements, initialOutline, importantInfo, nil
}

func (g *Generator) CountChapters(chapterOutline string) (uint, error) {
	prompt := fmt.Sprintf(prompts.CHAPTER_COUNT_PROMPT, chapterOutline)
	modelInfo, err := GetModelAndProvider(g.config.Models.ChapterOutline)
	if err != nil {
		return 0, err
	}

	messages := []RequestMessage{BuildMessage(prompt)}
	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return 0, err
	}

	output := GetLastMessage(messages)
	var data struct {
		TotalChapters int `json:"totalChapters"`
	}

	err = json.Unmarshal([]byte(output), &data)
	if err != nil {
		return 0, fmt.Errorf("unable to unmarshal chapter count: %w", err)
	}

	return uint(data.TotalChapters), nil
}

func (g *Generator) GenerateChapterOutlines(outline string, outlineCount uint) ([]string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.ChapterOutline)
	if err != nil {
		return nil, err
	}

	prompt := fmt.Sprintf("Please help me expand upon the following outline, chapter by chapter.\n%s", outline)
	messages := []RequestMessage{BuildMessage(prompt)}
	chapterOutlines := make([]string, 0, outlineCount)

	for i := range outlineCount {
		chapterNum := i + 1
		prompt := fmt.Sprintf(prompts.CHAPTER_OUTLINE_PROMPT, chapterNum, outline, chapterNum)
		messages = append(messages, BuildMessage(prompt))

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate chapter %d outline: %w", chapterNum, err)
		}

		chapterOutlines = append(chapterOutlines, GetLastMessage(messages))
	}

	return chapterOutlines, nil
}

func (g *Generator) GenerateChapters(chapterOutlines []string, chapterCount uint, storyOutline string, importantInfo string) ([]string, error) {
	chapters := make([]string, 0, chapterCount)
	modelInfo, err := GetModelAndProvider(g.config.Models.ChapterWriter)
	if err != nil {
		return nil, err
	}

	for i := range chapterCount {
		chapterNum := i + 1
		currentChapterOutline := chapterOutlines[i]

		// Build previous chapters context
		previousChapters := ""
		for j := range i {
			previousChapters += chapters[j] + "\n"
		}

		// Generate summary of last chapter
		lastChapterSummary := ""
		if i > 0 {
			prompt := fmt.Sprintf(prompts.CHAPTER_SUMMARY_PROMPT, chapterNum, storyOutline, chapters[i-1])
			messages := []RequestMessage{
				{Role: "system", Content: prompts.CHAPTER_SUMMARY_INTRO},
				BuildMessage(prompt),
			}

			messages, err = g.generateText(messages, modelInfo.Model)
			if err != nil {
				return nil, fmt.Errorf("failed to generate chapter %d summary: %w", chapterNum, err)
			}
			lastChapterSummary = GetLastMessage(messages)
		}

		previousChaptersContext := ""
		if previousChapters != "" {
			previousChaptersContext = fmt.Sprintf("Previous chapters context:\n%s", previousChapters)
		}

		// Stage 1: Generate Initial Plot
		prompt := fmt.Sprintf(prompts.CHAPTER_GENERATION_STAGE1,
			previousChaptersContext, importantInfo, chapterNum, chapterCount,
			currentChapterOutline, lastChapterSummary, chapterNum)

		messages := []RequestMessage{
			{Role: "system", Content: fmt.Sprintf(prompts.CHAPTER_GENERATION_INTRO, storyOutline)},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate chapter %d stage 1: %w", chapterNum, err)
		}
		firstChapterStage := GetLastMessage(messages)

		// Stage 2: Add Character Development
		prompt = fmt.Sprintf(prompts.CHAPTER_GENERATION_STAGE2,
			previousChaptersContext, importantInfo, chapterNum, chapterCount,
			currentChapterOutline, lastChapterSummary, firstChapterStage, chapterNum)

		messages = []RequestMessage{
			{Role: "system", Content: fmt.Sprintf(prompts.CHAPTER_GENERATION_INTRO, storyOutline)},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate chapter %d stage 2: %w", chapterNum, err)
		}
		secondChapterStage := GetLastMessage(messages)

		// Stage 3: Add Dialogue
		prompt = fmt.Sprintf(prompts.CHAPTER_GENERATION_STAGE3,
			previousChaptersContext, importantInfo, chapterNum, chapterCount,
			lastChapterSummary, secondChapterStage, chapterNum)

		messages = []RequestMessage{
			{Role: "system", Content: fmt.Sprintf(prompts.CHAPTER_GENERATION_INTRO, storyOutline)},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to generate chapter %d stage 3: %w", chapterNum, err)
		}
		thirdChapterStage := GetLastMessage(messages)

		chapters = append(chapters, thirdChapterStage)
	}

	return chapters, nil
}

func (g *Generator) ReviewOutline(storyOutline string) (string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.Revision)
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf(prompts.CRITIC_OUTLINE_PROMPT, storyOutline)
	messages := []RequestMessage{
		{Role: "system", Content: prompts.CRITIC_OUTLINE_INTRO},
		BuildMessage(prompt),
	}

	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", err
	}
	feedback := GetLastMessage(messages)

	prompt = fmt.Sprintf(prompts.OUTLINE_REVISION_PROMPT, storyOutline, feedback)
	messages = []RequestMessage{BuildMessage(prompt)}

	messages, err = g.generateText(messages, modelInfo.Model)
	if err != nil {
		return "", err
	}

	return GetLastMessage(messages), nil
}

func (g *Generator) ReviewChapters(chapters []string) ([]string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.Revision)
	if err != nil {
		return nil, err
	}

	reviewedChapters := make([]string, 0, len(chapters))
	for i, chapter := range chapters {
		prompt := fmt.Sprintf(prompts.CRITIC_CHAPTER_PROMPT, chapter)
		messages := []RequestMessage{
			{Role: "system", Content: prompts.CRITIC_CHAPTER_INTRO},
			BuildMessage(prompt),
		}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to get feedback for chapter %d: %w", i+1, err)
		}
		feedback := GetLastMessage(messages)

		prompt = fmt.Sprintf(prompts.CHAPTER_REVISION_PROMPT, chapter, feedback)
		messages = []RequestMessage{BuildMessage(prompt)}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to revise chapter %d: %w", i+1, err)
		}

		reviewedChapters = append(reviewedChapters, GetLastMessage(messages))
	}

	return reviewedChapters, nil
}

func (g *Generator) ScrubChapters(chapters []string) ([]string, error) {
	modelInfo, err := GetModelAndProvider(g.config.Models.Scrub)
	if err != nil {
		return nil, err
	}

	scrubbedChapters := make([]string, 0, len(chapters))
	for i, chapter := range chapters {
		prompt := fmt.Sprintf(prompts.SCRUB_PROMPT, chapter)
		messages := []RequestMessage{BuildMessage(prompt)}

		messages, err = g.generateText(messages, modelInfo.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to scrub chapter %d: %w", i+1, err)
		}

		scrubbedChapters = append(scrubbedChapters, GetLastMessage(messages))
	}

	return scrubbedChapters, nil
}

func (g *Generator) generateText(messages []RequestMessage, model string) ([]RequestMessage, error) {

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
