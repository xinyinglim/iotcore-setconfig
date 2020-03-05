package iotcoresetconfig

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

//StepperConfig has one arg
//NewPos "A","B", "C"
type StepperConfig struct {
	NewPos string `json:"new_pos"`
}

// Stepper Class for configuring Stepper motors.
// Switches between several positions
// For this example, there are 3 possible positions: A, B and C dependent on the weather
type Stepper struct {
	DeviceID string
	*StepperConfig
}

func (s *Stepper) GetDeviceID() string {
	return s.DeviceID
}

func (s *Stepper) GetPath() string {
	return fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", projectID, location, stepperRegistryID, s.DeviceID)
}

func (s *Stepper) GetConfigJSON() (string, error) {
	if s.NewPos == "" {
		return "{}", fmt.Errorf("Not configured")
	}
	bytes, err := json.Marshal((*(s.StepperConfig)))
	if err != nil {
		return "", err
	}
	return b64.StdEncoding.EncodeToString(bytes), nil
}

func (s *Stepper) GetConfigDetails() string {
	return fmt.Sprintf("Stepper has new position of %s\n", s.NewPos)
}

func (s *Stepper) UpdateConfig(newWeather Weather) {
	s.StepperConfig = &StepperConfig{}
	switch newWeather {
	case Sunny:
		s.NewPos = "A"
	case Rainy:
		s.NewPos = "B"
	case Cloudy:
		s.NewPos = "C"
	}
}
