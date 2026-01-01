package console

import (
	"fmt"
	objects "shell-aid/lib"
	"strings"
)

const (
	Width uint8 = 180

	Reset string = "\033[0m"
	Cyan  string = "\033[36m"
	Red   string = "\033[31m"
	Green string = "\033[32m"
	Bold  string = "\033[1m"
)

func wrapText(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	var result strings.Builder
	lineLength := 0

	for i, word := range words {
		wordLen := len(word)

		if lineLength+wordLen+1 > int(Width) {
			result.WriteString("\n")
			lineLength = 0
		}

		if i > 0 && lineLength > 0 {
			result.WriteString(" ")
			lineLength++
		}

		result.WriteString(word)
		lineLength += wordLen
	}

	return result.String()
}

func Display(resp objects.GenerateCommandResponse) {
	fmt.Println()
	fmt.Println(Cyan + Bold + "WORKFLOW" + Reset)
	fmt.Println(wrapText(resp.Workflow))
	fmt.Println()

	fmt.Println(Red + Bold + "SIDE EFFECTS" + Reset)
	fmt.Println(wrapText(resp.SideEffects))
	fmt.Println()

	fmt.Println(Green + Bold + "COMMAND" + Reset)
	fmt.Println(resp.Command)
	fmt.Println()

}
