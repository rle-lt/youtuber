# Large context story generation

A Go library and CLI tool for generating fictional stories using AI models.

## Installation

### As a CLI tool

```bash
go install scripter/cmd/scripter@latest
```

### As a library

```bash
go get scripter/pkg/scripter
```

## Usage

### Command Line

```bash
# Set your API key
export OPENROUTER_API_KEY="your-api-key"

# Generate a story
scripter -prompt /path/to/prompt.txt > story.txt
```

### As a Library

```go
package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    
    "scripter/internal/constants"
    "scripter/pkg/scripter"
)

func main() {
    // Configure the generator
    config := scripter.Config{
        APIKey: os.Getenv("OPENROUTER_API_KEY"),
        Models: scripter.Models{
            InitialOutline:  "openrouter://amazon/nova-2-lite-v1:free",
            ChapterOutline:  "openrouter://amazon/nova-2-lite-v1:free",
            ChapterWriter:   "openrouter://amazon/nova-2-lite-v1:free",
            Revision:        "openrouter://amazon/nova-2-lite-v1:free",
            Scrub:           "openrouter://amazon/nova-2-lite-v1:free",
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
```

### Using Individual Components

```go
package main

import (
    "log"
    "scripter/pkg/scripter"
)

func main() {
    config := scripter.Config{
        APIKey: "your-api-key",
        Models: scripter.Models{
            InitialOutline: "openrouter://amazon/nova-2-lite-v1:free",
            // ... other models
        },
        MaxChapterCount: 6,
    }

    gen, err := scripter.NewGenerator(config)
    if err != nil {
        log.Fatal(err)
    }

    // Just generate an outline
    outline, elements, _, info, err := gen.GenerateInitOutline("Your prompt here")
    if err != nil {
        log.Fatal(err)
    }

    // Or count chapters
    count, err := gen.CountChapters(outline)
    if err != nil {
        log.Fatal(err)
    }
    
    // etc...
}
```

### Environment Variables

- `OPENROUTER_API_KEY`: Your OpenRouter API key (required)

## Building

```bash
# Build the CLI
go build -o scripter ./cmd/scripter

# Run tests (if you add them)
go test ./...
```
