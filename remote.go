package main

import (
	"encoding/json"
	"fmt"
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

func CurrentSong() (Song, error) {
	jsonData, err := jsonRequest("now-playing")
	if err != nil {
		return Song{}, err
	}

	var data struct {
		Song Song `json:"info"`
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return Song{}, err
	}

	return data.Song, nil
}

/*

	POST requests

*/

func PlayURL(url string) error {
	data := map[string]string{"url": url}
	jsonData, _ := json.Marshal(data)

	return postRequest("play-url", jsonData)
}

func PlayWithHref(href string) error {
	data := map[string]string{"href": href}
	jsonData, _ := json.Marshal(data)

	return postRequest("play-item-href", jsonData)
}

func PlayItem(itemType ItemType, id string) error {
	data := map[string]string{
		"type": string(itemType),
		"id":   id,
	}
	jsonData, _ := json.Marshal(data)

	return postRequest("play-item", jsonData)
}

func JumpTo(seconds int) error {
	data := map[string]int{"position": seconds}
	jsonData, _ := json.Marshal(data)

	return postRequest("seek", jsonData)
}

func AddCurrentSongToLibrary() error {
	return postRequestNoJson("add-to-library")
}

/*

	TESTING

*/

// testing only
func main() {
	fmt.Println(AddCurrentSongToLibrary())
}
