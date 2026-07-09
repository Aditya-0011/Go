package api

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func GetCoordinates(city string) (io.Reader, error) {
	apiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	if apiKey == "" {
		return nil, errors.New("OPEN_WEATHER_MAP_API_KEY not found")
	}

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, apiKey)
	return fetch(url)
}
