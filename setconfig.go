package iotcoresetconfig

//todo open source this and make it into a tutorial: this is non intuitive
//gcloud functions deploy UpdateWeather --runtime go111 --region asia-east2 --trigger-http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	cloudiot "google.golang.org/api/cloudiot/v1"
)

func SetWeather(newWeather Weather) error {
	ctx := context.Background()
	service, err := cloudiot.NewService(ctx)
	if err != nil {
		return err
	}

	devices := GetDevices()
	deviceService := cloudiot.NewProjectsLocationsRegistriesDevicesService(service)

	for _, d := range devices {
		d.UpdateConfig(newWeather)
		// fmt.Printf(d.GetConfigDetails())
		encodedString, err := d.GetConfigJSON()
		if err != nil {
			fmt.Printf("Failed to update %s : %v\n", d.GetDeviceID(), err)
			continue
		}
		// fmt.Println(encodedString)
		configRequest := cloudiot.ModifyCloudToDeviceConfigRequest{
			BinaryData: encodedString,
		}
		call := deviceService.ModifyCloudToDeviceConfig(d.GetPath(), &configRequest)
		call.Context(ctx)

		_, err = call.Do()
		fmt.Printf("called %s", d.GetDeviceID())
		if err != nil {
			fmt.Printf("Failed to update %s : %v\n", d.GetDeviceID(), err)
			continue
		}
	}

	return nil

}

type WeatherFS struct {
	CurrentWeather Weather   `json:"currentWeather"`
	LastUpdated    time.Time `json:"lastUpdated"`
}

type WeatherUpdate struct {
	Weather Weather `json:"weather"`
}

func UpdateWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		//preflight request
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	receivedTime := time.Now()
	var errString string

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var weatherUpdate WeatherUpdate
		err := json.NewDecoder(r.Body).Decode(&weatherUpdate)
		fmt.Println("succesfully gotten weather update")
		if err != nil {
			errString = fmt.Sprintf("Error in passing app: %v", err)
			break
		}
		//todo check firestore, if is new update, then update and firestore and continue
		//else it is a repeat
		var firestoreClient *firestore.Client
		ctx := context.Background()
		firestoreClient, err = firestore.NewClient(ctx, projectID)
		if err != nil {
			errString = fmt.Sprintf("Error: %v", err)
			fmt.Println(errString)
			break
		}

		err = firestoreClient.RunTransaction(ctx, func(ctxt context.Context, tx *firestore.Transaction) error {
			docRef := firestoreClient.Collection("currentWeather").Doc("currentWeather")
			newWeather := weatherUpdate.Weather
			snap, e := tx.Get(docRef)
			if err != nil {
				return err
			}
			var weatherFS WeatherFS
			snap.DataTo(&weatherFS)
			if receivedTime.After(weatherFS.LastUpdated) {
				e = tx.Update(docRef, []firestore.Update{{
					Path:  "currentWeather",
					Value: newWeather,
				}, {
					Path:  "lastUpdated",
					Value: receivedTime,
				}})
				if e != nil {
					return e
				}
				return SetWeather(newWeather)
			}
			return nil

		})

		if err != nil {
			errString = fmt.Sprintf("Error: %v", err)
			fmt.Println(errString)
			break
		}

	}
	var result Result
	if errString != "" {
		result = Result{IsSuccessful: false, ErrorMessage: errString}
	} else {
		result = Result{IsSuccessful: true}
	}
	bytes, _ := json.Marshal(result)

	w.Header().Set("Context-Type", "application/json")
	w.Write(bytes)
}

type Result struct {
	IsSuccessful bool   `json:"isSuccessful"`
	ErrorMessage string `json:"errorMessage"`
}
