package cidergo

import "encoding/json"

func AddItemToQueue(itemType ItemType, id string, last bool) error {
	data := map[string]string{
		"type": string(itemType),
		"id":   id,
	}
	jsonData, _ := json.Marshal(data)

	if last {
		return postRequest("play-later", jsonData)
	} else {
		return postRequest("play-next", jsonData)
	}
}

func MoveQueueItem(startIndex int, destinationIndex int) error {
	data := map[string]int{
		"startIndex":       startIndex,
		"destinationIndex": destinationIndex,
	}
	jsonData, _ := json.Marshal(data)

	return postRequest("/queue/move-to-position", jsonData)
}

func DeleteQueueItem(index int) error {
	data := map[string]int{"index": index}
	jsonData, _ := json.Marshal(data)

	return postRequest("/queue/remove-by-index", jsonData)
}

func clearQueue() error {
	return postRequestNoJson("queue/clear-queue")
}

func GetQueue() ([]QueueItem, error) {
	jsonData, err := jsonRequest("queue")
	if err != nil {
		return []QueueItem{}, err
	}

	var data []QueueItem
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return []QueueItem{}, err
	}

	return data, nil
}
