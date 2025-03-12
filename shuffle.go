package cidergo

import "encoding/json"

type ShuffleMode int

const (
	ShuffleModeOff ShuffleMode = 0
	ShuffleModeOn  ShuffleMode = 1
)

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

func SetShuffleMode(mode ShuffleMode) error {
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

func ToggleShuffle() error {
	return postRequestNoJson("toggle-shuffle")
}
