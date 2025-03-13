package cidergo

import "encoding/json"

// PlayURL plays and item using its APPLE MUSIC link (not Cider link or song link)
func PlayURL(url string) error {
	data := map[string]string{"url": url}
	jsonData, _ := json.Marshal(data)

	return postRequest("play-url", jsonData)
}

// PlayWithHref plays an item using its Href (Apple Music API identifier)
func PlayWithHref(href string) error {
	data := map[string]string{"href": href}
	jsonData, _ := json.Marshal(data)

	return postRequest("play-item-href", jsonData)
}

// PlayItem plays an item using its ID (for example, Song.PlayParams.ID). ItemType has to be specified!
func PlayItem(itemType ItemType, id string) error {
	data := map[string]string{
		"type": string(itemType),
		"id":   id,
	}
	jsonData, _ := json.Marshal(data)

	return postRequest("play-item", jsonData)
}
