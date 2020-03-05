package iotcoresetconfig

import (
	"testing"
)

func TestSetWeather(t *testing.T) {
	err := SetWeather(Rainy)
	if err != nil {
		t.Errorf("Error of weather: %v", err)
	}
	t.Logf("Successful set weather")
}
