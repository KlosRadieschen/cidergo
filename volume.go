package cidergo

import (
	"encoding/json"
	"errors"
	"time"
)

// GetVolume returns the current volume of the client
func GetVolume() (float64, error) {
	data := make(map[string]any)

	// Only waits if it is NOT the first run
	for firstRun := true; data["value"] == nil; firstRun = false {
		if !firstRun {
			time.Sleep(waitDelay)
		}

		jsonData, err := jsonRequest("volume")
		if err != nil {
			return -1, err
		}

		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			return -1, err
		}
	}

	return data["volume"].(float64), nil
}

// SetVolume changes the volume of the client to the value passed as parameter (must be between 0 and 1)
func SetVolume(volume float64) error {
	if volume < 0 || volume > 1 {
		return errors.New("volume must be between 0 and 1")
	}

	data := map[string]float64{"volume": volume}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}

// AdjustVolume adds or subtracts (if negative) the value passed as parameter from the current volume
// For example, if the current volume is 0.7, AdjustVolume(-0.5) will set it to 0.2
func AdjustVolume(amount float64) error {
	currentVolume, err := GetVolume()
	if err != nil {
		return err
	}

	data := map[string]float64{"volume": currentVolume + amount}
	jsonData, _ := json.Marshal(data)

	err = postRequest("volume", jsonData)
	if err != nil && err.Error() == "HTTP error: 500 Internal Server Error" {
		return errors.New("volume out of bounds (must be between 0 and 1)")
	}
	return postRequest("volume", jsonData)
}
