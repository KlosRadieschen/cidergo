package cidergo

import "encoding/json"

// Like tells Apple Music that you like the currently playing song
func Like() error {
	data := map[string]int{"rating": 1}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}

// Dislike tells Apple Music that you dislike the currently playing song
func Dislike() error {
	data := map[string]int{"rating": -1}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}

// ResetLike removes your like or dislike from the currently playing song
func ResetLike() error {
	data := map[string]int{"rating": 0}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}
