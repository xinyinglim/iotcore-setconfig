package iotcoresetconfig

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

type DCMotor struct {
	DeviceID string
	*DCMotorConfig
}

func (s *DCMotor) UpdateConfig(newWeather Weather) {
	s.DCMotorConfig = &DCMotorConfig{}
	switch newWeather {
	case Sunny:
		s.DCMotorConfig.On = true
		s.DCMotorConfig.Speed = 40
	case Rainy:
		s.DCMotorConfig.On = false
		s.DCMotorConfig.Speed = 0
	case Cloudy:
		s.DCMotorConfig.On = true
		s.DCMotorConfig.Speed = 15
	}
}

func (s *DCMotor) GetDeviceID() string {
	return s.DeviceID
}

func (s *DCMotor) GetPath() string {
	return fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", projectID, location, dcmotorRegistryID, s.DeviceID)
}

func (s *DCMotor) GetConfigDetails() string {
	return fmt.Sprintf("DCMotor: on: %t with speed %d\n", s.On, s.Speed)
}

func (s *DCMotor) GetConfigJSON() (string, error) {
	bytes, err := json.Marshal(*(s.DCMotorConfig))
	if err != nil {
		return "", err
	}
	return b64.StdEncoding.EncodeToString(bytes), nil
}

//DCMotorConfig has two args
//MotionType is "wave" or "move_pos"
type DCMotorConfig struct {
	On    bool `json:"on"`
	Speed int  `json:"speed"`
}
