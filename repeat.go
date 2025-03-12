package main

import (
	"encoding/json"
)

type RepeatMode int

const (
	RepeatModeOff  RepeatMode = 0
	RepeatModeSong RepeatMode = 1
	RepeatModeOn   RepeatMode = 2
)

func GetRepeat() (RepeatMode, error) {
	jsonData, err := jsonRequest("repeat-mode")
	if err != nil {
		return -1, err
	}

	data := make(map[string]any)
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return -1, err
	}

	return RepeatMode(data["value"].(float64)), nil
}

func SetRepeat(option RepeatMode) error {
	currentRepeat, err := GetRepeat()
	if err != nil {
		return err
	}

	if option < currentRepeat {
		option += 3
	}
	difference := int(option - currentRepeat)

	for i := 0; i < difference; i++ {
		if err := ToggleRepeat(); err != nil {
			return err
		}
	}

	return nil
}

func ToggleRepeat() error {
	return postRequestNoJson("toggle-repeat")
}
