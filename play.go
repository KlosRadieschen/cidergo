package cidergo

import "encoding/json"

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
