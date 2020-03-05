# Google Cloud Functions to IoT Core

This is part of the demo IoT Cafe.

Upon triggering a new weather pattern on the dashboard, Google Cloud Functions is used to update IoT Core devices configurations.

Examples for the following 3 types of devices are given
- Servo Motor
- Stepper Motor
- DC Motor

# To use
1. Set up IoT Core in GCP and create and add relevant devices (Servo, stepper or DC motor)
2. In config.go, add registry ID
3. In devices.go, add list of all devices, their types and device name in IoT Core to GetDevices function
4. Deploy GCP Cloud Function with command line
gcloud functions deploy UpdateWeather --runtime go111 --region \[GCP-REGION\] --trigger-http
