package iotcoresetconfig

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

// Servo holds data for Servo Motor
// 2 types of Motion: Waving and Moving to a specific position
type Servo struct {
	DeviceID string
	*ServoConfig
}

func (s *Servo) GetDeviceID() string {
	return s.DeviceID
}

func (s *Servo) GetPath() string {
	return fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", projectID, location, servoRegistryID, s.DeviceID)
}

func (s *Servo) GetConfigJSON() (string, error) {
	if s.MotionType == "" {
		return "{}", fmt.Errorf("Not configured")
	}
	bytes, err := json.Marshal(*(s.ServoConfig))
	if err != nil {
		return "", err
	}
	return b64.StdEncoding.EncodeToString(bytes), nil
}

func (s *Servo) GetConfigDetails() string {
	return fmt.Sprintf("Servo is motiontype: %s, position %d\n", s.MotionType, s.Pos)
}

//ServoConfig has two args
//MotionType is "wave" or "move_pos"
type ServoConfig struct {
	MotionType string `json:"motion_type"`
	Pos        int    `json:"pos,omitempty"`
}

func (s *Servo) UpdateConfig(newWeather Weather) {
	s.ServoConfig = &ServoConfig{}
	switch newWeather {
	case Sunny:
		// s.ServoConfig.MotionType = "wave"
		s.ServoConfig.MotionType = "move_pos"
		s.ServoConfig.Pos = 1
	case Rainy:
		s.ServoConfig.MotionType = "move_pos"
		s.ServoConfig.Pos = 1
	case Cloudy:
		s.ServoConfig.MotionType = "move_pos"
		s.ServoConfig.Pos = -1
	}
}
