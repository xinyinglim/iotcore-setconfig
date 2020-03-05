package iotcoresetconfig

import "fmt"

//todo generalize this
type Weather string

const (
	Sunny  Weather = "sunny"
	Rainy  Weather = "rainy"
	Cloudy Weather = "cloudy"
)

var messageToWeather = map[string]Weather{
	"sunny":  Sunny,
	"rainy":  Rainy,
	"cloudy": Cloudy,
}

func MessageToWeather(message string) (Weather, error) {
	weather, ok := messageToWeather[message]
	if !ok {
		return "", fmt.Errorf("New weather message: %s", message)
	}
	return weather, nil
}

type CurrentWeather struct {
	CurrentWeather Weather `firestore:"currentWeather",json:"currentWeather"`
}
