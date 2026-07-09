package api

import (
	"fmt"
	"io"
)

func GetWeather(lat, long string) (io.Reader, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current=temperature_2m,relative_humidity_2m,precipitation,rain,showers,apparent_temperature,is_day,cloud_cover,pressure_msl,surface_pressure,wind_speed_10m,wind_direction_10m,wind_gusts_10m", lat, long)
	return fetch(url)
}
