package cidergo

import (
	"encoding/json"
	"fmt"
	"time"
)

// waitDelay represents the amount of time (in millisecond) to wait before trying again when a GET request return OK without an actual value.
// This happens when you use a Toggle/Set function for a feature like shuffle and then instantly use the Get function for that same feature
var waitDelay = 300 * time.Millisecond

// CheckConnection tries to establish a connection with the Cider RPC. Returns nil when successful
func CheckConnection() error {
	_, err := statusRequest("active")
	return err
}

// CheckPlaying returns a boolean telling you if Cider is currently playing something
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

// JumpTo jumps to a certain time in the song. Only whole seconds are accepted (Be careful, as Song holds time in milliseconds)
func JumpTo(seconds int) error {
	data := map[string]int{"position": seconds}
	jsonData, _ := json.Marshal(data)

	return postRequest("seek", jsonData)
}

// JumpToPercentage jumps to the moment in the song where x% have passed (where x is the percentage parameter)
func JumpToPercentage(percentage int) error {
	if percentage < 0 || percentage > 100 {
		return fmt.Errorf("invalid jump percentage: %d", percentage)
	}

	currentSong, err := CurrentSong()
	if err != nil {
		return err
	}
	seconds := ((currentSong.DurationInMillis / 1000) / 100) * percentage // I know this could be simplified, but I think the steps are more readable this way

	data := map[string]int{"position": seconds}
	jsonData, _ := json.Marshal(data)

	return postRequest("seek", jsonData)
}

// AddCurrentSongToLibrary will add the currently playing song to your Apple Music library
func AddCurrentSongToLibrary() error {
	return postRequestNoJson("add-to-library")
}

// SetWaitDelay changes the waitDelay (default: 300), which represents the amount of time (in millisecond) to wait before trying again when a GET request returns OK without an actual value.
// This happens when you use a Toggle/Set function for a feature like shuffle and then instantly use the Get function for that same feature.
// Error only occurs when passing a negative number
func SetWaitDelay(millis int) error {
	if millis < 0 {
		return fmt.Errorf("wait delay cannot be negative: %d", millis)
	}
	waitDelay = time.Duration(millis) * time.Millisecond
	return nil
}
