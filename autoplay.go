package main

import (
	"encoding/json"
)

type AutoplayMode bool

const (
	AutoplayModeOff AutoplayMode = false
	AutoplayModeOn  AutoplayMode = true
)

func GetAutoplay() (AutoplayMode, error) {
	jsonData, err := jsonRequest("autoplay")
	if err != nil {
		return false, err
	}

	data := make(map[string]any)
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return false, err
	}

	return AutoplayMode(data["value"].(bool)), nil
}

func SetAutoplayMode(mode AutoplayMode) error {
	currentAutoplay, err := GetAutoplay()
	if err != nil {
		return err
	}

	if currentAutoplay != mode {
		if err := ToggleAutoplay(); err != nil {
			return err
		}
	}
	return nil
}

func ToggleAutoplay() error {
	return postRequestNoJson("toggle-autoplay")
}
