package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Remote struct {
	port string
}

func (r *Remote) CheckConnection() error {
	_, err := statusRequest(r.port, "active")

	return err
}

func (r *Remote) CurrentSong() (Song, error) {
	jsonData, err := jsonRequest(r.port, "now-playing")

	var data struct {
		Song Song `json:"info"`
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return Song{}, err
	}

	return data.Song, nil
}

func Connect(port string) (*Remote, error) {
	if port == "" {
		port = "10767"
	}

	r := Remote{port: port}

	return &r, r.CheckConnection()
}

func statusRequest(port string, endpoint string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("http://0.0.0.0:%s/api/v1/playback/%s", port, endpoint))
	if err != nil {
		return 0, err
	} else if resp.StatusCode == 404 {
		return 0, errors.New(fmt.Sprintf("endpoint /%s not found", endpoint))
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func jsonRequest(port string, endpoint string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://0.0.0.0:%s/api/v1/playback/%s", port, endpoint))
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode == 404 {
		return []byte{}, errors.New(fmt.Sprintf("endpoint /%s not found", endpoint))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return body, nil
}

func main() {
	r, err := Connect("")
	if err != nil {
		panic(err)
	}

	song, err := r.CurrentSong()
	if err != nil {
		panic(err)
	}

	fmt.Println(song)
}
