package main

import (
	"encoding/json"
	"errors"
)

func GetVolume() (float64, error) {
	jsonData, err := jsonRequest("volume")
	if err != nil {
		return -1, err
	}

	data := make(map[string]any)
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return -1, err
	}

	return data["volume"].(float64), nil
}

func SetVolume(volume float64) error {
	data := map[string]float64{"volume": volume}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}

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
