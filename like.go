package main

import "encoding/json"

func Like() error {
	data := map[string]int{"rating": 1}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}

func Dislike() error {
	data := map[string]int{"rating": -1}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}

func ResetLike() error {
	data := map[string]int{"rating": 0}
	jsonData, _ := json.Marshal(data)

	return postRequest("volume", jsonData)
}
