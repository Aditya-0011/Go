package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	objects "shell-aid/lib"
	"shell-aid/utils/ai"
	"shell-aid/utils/console"
	"shell-aid/utils/shell"
	"strings"
)

func main() {
	sysInfo := shell.Detect()

	agent := true

	for agent {
		fmt.Println("Enter your request (press Enter on an empty line to finish), enter 'exit' to quit:")

		var b strings.Builder
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			line := scanner.Text()

			if line == "exit" {
				fmt.Println("Exiting.")
				os.Exit(0)
			}

			if line == "" {
				fmt.Println("Processing Request")
				break
			}

			b.WriteString(line)
			b.WriteString("\n")
		}

		if err := scanner.Err(); err != nil {
			log.Fatalln("Error reading from stdin:", err)
		}

		if b.Len() == 0 {
			fmt.Println("Input is required")
			continue
		}

		resp := ai.GenerateCommand(objects.GenerateCommandRequest{Command: b.String(), SysInfo: sysInfo})

		var cmdResp objects.GenerateCommandResponse
		if err := json.Unmarshal([]byte(resp), &cmdResp); err != nil {
			fmt.Println(resp)
		} else {
			fmt.Println(b.String())
			console.Display(cmdResp)

		}
	}
}
