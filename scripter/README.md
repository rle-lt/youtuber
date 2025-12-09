# Large context story generation

Go module for large story generation

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
export OPENROUTER_API_KEY="your-api-key"

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

    generator, err := scripter.NewGenerator(config)
    if err != nil {
        log.Fatal(err)
    }

    prompt := "Write a science fiction story about a lone astronaut..."
    chapters, err := generator.GenerateStory(prompt)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(strings.Join(chapters, "\n\n"))
}
```

### Generating individual story components

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
            // ... 
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

    ...
}
```

## Building

```bash
go build -o scripter ./cmd/scripter
``
