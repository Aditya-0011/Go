package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	objects "shell-aid/lib"
	"strings"

	"google.golang.org/genai"
)

const key string = ""

func sanitizeJSON(input string) string {
	input = strings.TrimSpace(input)

	if strings.HasPrefix(input, "```") {
		if i := strings.Index(input, "\n"); i != -1 {
			input = input[i+1:]
		}
		if j := strings.LastIndex(input, "```"); j != -1 {
			input = input[:j]
		}
		input = strings.TrimSpace(input)
	}

	if !json.Valid([]byte(input)) {
		if i := strings.Index(input, "{"); i != -1 {
			if j := strings.LastIndex(input, "}"); j != -1 && j >= i {
				candidate := input[i : j+1]
				if json.Valid([]byte(candidate)) {
					return candidate
				}
			}
		}
	}

	return input
}

func GenerateCommand(props objects.GenerateCommandRequest) string {
	cmd := props.Command
	os := props.SysInfo.OS
	shell := props.SysInfo.Shell

	ctx := context.Background()
	defer ctx.Done()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: key})
	if err != nil {
		log.Fatal(err)
	}

	sysInstruction := genai.Text(fmt.Sprintf(`
	You are a shell expert for OS: %s and its shell: %s. You have to generate shell commands based on user requests. Output will be json object with have three fields: "Workflow" (string), "SideEffects" (string) and "Command" (string). If the user has a request which requires multiple commands to be executed, then you will chain them according to the shell syntax. Workflow should be a step by step explanation of what the command does. SideEffects should list any potential side effects of running the command. Command should be the exact command to run. Both SideEffect and Workflow must be in paragraph(s).
	`, os, shell))

	config := &genai.GenerateContentConfig{SystemInstruction: sysInstruction[0]}

	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(cmd), config)
	if err != nil {
		log.Fatal(err)
	}
	return sanitizeJSON(result.Text())
}
