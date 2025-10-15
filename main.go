package main

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"io"
	"log"
	"os"
)

func main() {
	var input string
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("Failed to get stdin info: %v", err)
	}

	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	isPiped := (stat.Mode() & os.ModeCharDevice) == 0
	if isPiped {
		piped, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading from stdin: %v", err)
		}
		input += string(piped) + " "
	}

	if len(os.Args) > 1 {
		input += os.Args[1]
	}

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
