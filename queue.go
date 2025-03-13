package cidergo

import "encoding/json"

// AddItemToQueue adds and item to the queue, using its ID (for example, Song.PlayParams.ID). ItemType has to be specified!
// last indicates whether the item is added to the end (false) or start (true) of the queue
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

// MoveQueueItem moves an item in the queue from one index to another
func MoveQueueItem(startIndex int, destinationIndex int) error {
	data := map[string]int{
		"startIndex":       startIndex,
		"destinationIndex": destinationIndex,
	}
	jsonData, _ := json.Marshal(data)

	return postRequest("/queue/move-to-position", jsonData)
}

// DeleteQueueItem deletes a single item from the queue using its index
func DeleteQueueItem(index int) error {
	data := map[string]int{"index": index}
	jsonData, _ := json.Marshal(data)

	return postRequest("/queue/remove-by-index", jsonData)
}

// ClearQueue deletes all items from the queue
func ClearQueue() error {
	return postRequestNoJson("queue/clear-queue")
}

// GetFullQueue returns the entire queue with all the relevant metadata. THE CURRENTLY PLAYING QueueItem IS NOT NECESSARILY AT INDEX 0 (can contain part of the history)
func GetFullQueue() ([]QueueItem, error) {
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
