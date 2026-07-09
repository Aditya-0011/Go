package tools

import (
	"agent/api"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/tool"
)

func Coordinates() tool.Tool {
	tool, err := create("get_coordinates", "Retrieves a list of possible coordinates (latitude and longitude) for a specified city. Review the results (including country and state) and select the most appropriate one.",
		func(ctx agent.Context, args cityArgs) (map[string]any, error) {
			coordsReader, err := api.GetCoordinates(args.City)
			if err != nil {
				return map[string]any{
					"status":        "error",
					"error_message": "Failed to get coordinates for '" + args.City + "': " + err.Error(),
				}, nil
			}

			var coords []map[string]any
			if err := json.NewDecoder(coordsReader).Decode(&coords); err != nil {
				return nil, fmt.Errorf("failed to decode coordinates: %w", err)
			}

			if len(coords) == 0 {
				return map[string]any{
					"status":        "error",
					"error_message": "Coordinates for '" + args.City + "' are not available.",
				}, nil
			}

			return map[string]any{
				"status":  "success",
				"results": coords,
			}, nil
		})
	if err != nil {
		log.Fatalf("Failed to create get_coordinates tool: %v", err)
		os.Exit(1)
	}

	return tool
}
