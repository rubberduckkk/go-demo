package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	item := "eat dinner with my friend Jack Bole"
	categories := "[Work, Heath, Travel]"
	prompt := fmt.Sprintf(`
Here are the available categories: %v. Pick the most suitable category for this todo list item: %v. If none is suitable, give a new one. Only return the category name. `, categories, item)
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
