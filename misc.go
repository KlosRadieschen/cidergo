package cidergo

import (
	"encoding/json"
)

func CheckConnection() error {
	_, err := statusRequest("active")
	return err
}

func isPlaying() (bool, error) {
	jsonData, err := jsonRequest("is-playing")
	if err != nil {
		return false, err
	}

	unmarshal := make(map[string]any)
	err = json.Unmarshal(jsonData, &unmarshal)
	if err != nil {
		return false, err
	}

	return unmarshal["is_playing"] == true, nil
}

func JumpTo(seconds int) error {
	data := map[string]int{"position": seconds}
	jsonData, _ := json.Marshal(data)

	return postRequest("seek", jsonData)
}

func AddCurrentSongToLibrary() error {
	return postRequestNoJson("add-to-library")
}
