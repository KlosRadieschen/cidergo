package cidergo

import "encoding/json"

// ShuffleMode represents the status of shuffle, as defined by Apple Music
type ShuffleMode int

const (
	ShuffleModeOff ShuffleMode = 0
	ShuffleModeOn  ShuffleMode = 1
)

// GetShuffle returns the current status of shuffle
func GetShuffle() (ShuffleMode, error) {
	jsonData, err := jsonRequest("shuffle-mode")
	if err != nil {
		return -1, err
	}

	data := make(map[string]any)
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return -1, err
	}

	return ShuffleMode(data["value"].(float64)), nil
}

// SetShuffle sets shuffle to the mode passes as parameter
func SetShuffle(mode ShuffleMode) error {
	currentShuffle, err := GetShuffle()
	if err != nil {
		return err
	}

	if currentShuffle != mode {
		if err := ToggleShuffle(); err != nil {
			return err
		}
	}
	return nil
}

// ToggleShuffle toggles the current state of shuffle
func ToggleShuffle() error {
	return postRequestNoJson("toggle-shuffle")
}
