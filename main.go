package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Printf(":>")
	line, _ := input.ReadString('\n')
	err, resp := prompt(line)

	fmt.Println(err, resp)
}

func prompt(str string) (error, string) {
	godotenv.Load()
	ctx := context.Background()
	apiKey := os.Getenv("KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		return err, ""
	}

	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, str)
	if err != nil {
		return err, ""
	}

	return nil, answer
}
