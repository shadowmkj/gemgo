package main

import (
	"os"
	"context"
	"fmt"
	"log"
	"google.golang.org/genai"
)

func main() {
	input := os.Args[1]
	if input == "" {
		log.Fatal("Please provide an input text as a command line argument.")
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)

	if err != nil {
		log.Fatalf("Failed to create client: %v. Make sure GEMINI_API_KEY is set.", err)
	}

	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		genai.Text(input),
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
}

