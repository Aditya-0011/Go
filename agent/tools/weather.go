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

func Weather() tool.Tool {
	tool, err := create("get_weather", "Retrieves the current weather report for specified coordinates (latitude and longitude).",
		func(ctx agent.Context, args coordinatesArgs) (map[string]any, error) {
			weatherReader, err := api.GetWeather(args.Lat, args.Lon)
			if err != nil {
				return map[string]any{
					"status":        "error",
					"error_message": "Failed to get weather for coordinates: " + err.Error(),
				}, nil
			}

			var weatherData map[string]any
			if err := json.NewDecoder(weatherReader).Decode(&weatherData); err != nil {
				return nil, fmt.Errorf("failed to decode weather data: %w", err)
			}

			return map[string]any{
				"status": "success",
				"report": weatherData,
			}, nil
		})
	if err != nil {
		log.Fatalf("Failed to create get_weather tool: %v", err)
		os.Exit(1)
	}

	return tool
}
