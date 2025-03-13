package cidergo

import (
	"encoding/json"
	"time"
)

// AutoplayMode represents the status of autoplay. Unlike Shuffle, it uses booleans instead of integers
type AutoplayMode bool

const (
	AutoplayModeOff AutoplayMode = false
	AutoplayModeOn  AutoplayMode = true
)

// GetAutoplay gets the current status of autoplay
func GetAutoplay() (AutoplayMode, error) {
	data := make(map[string]any)

	// Only waits if it is NOT the first run
	for firstRun := true; data["value"] == nil; firstRun = false {
		if !firstRun {
			time.Sleep(waitDelay)
		}

		jsonData, err := jsonRequest("autoplay")
		if err != nil {
			return false, err
		}

		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			return false, err
		}
	}

	return AutoplayMode(data["value"].(bool)), nil
}

// SetAutoplay sets autoplay to the mode passed as parameter
func SetAutoplay(mode AutoplayMode) error {
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

// ToggleAutoplay toggles the current state of autoplay
func ToggleAutoplay() error {
	return postRequestNoJson("toggle-autoplay")
}
