package iotcoresetconfig

type Device interface {
	GetPath() string
	GetConfigJSON() (string, error)
	UpdateConfig(newWeather Weather)
	GetDeviceID() string
	GetConfigDetails() string
}

//todo substitute the below with your own devices' id from IoT Core
func GetDevices() []Device {
	servo1 := &Servo{DeviceID: "SERVO-ID"}
	dcmotor1 := &DCMotor{DeviceID: "DCMOTOR-DEVICE-ID"}
	stepper1 := &Stepper{DeviceID: "STEPPER-DEVICE-ID"}
	return []Device{servo1, dcmotor1, stepper1}
}
