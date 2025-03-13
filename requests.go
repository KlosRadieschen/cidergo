package cidergo

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
)

const baseURL = "http://0.0.0.0:10767/api/v1/playback/"

// statusRequest performs a GET request to the endpoint passed as parameter and returns the HTML status as int
func statusRequest(endpoint string) (int, error) {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return 0, err
	} else if resp.StatusCode != 200 {
		return 0, errors.New("HTTP error: " + resp.Status)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

// jsonRequest performs a GET request to the endpoint passed as parameter and returns the JSON received from that request
func jsonRequest(endpoint string) ([]byte, error) {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode != 200 {
		return []byte{}, errors.New("HTTP error: " + resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return body, nil
}

// postRequest performs a POST request to the endpoint passed as parameter and sends the jsonData (in bytes) passed as parameter
func postRequest(endpoint string, jsonData []byte) error {
	req, err := http.NewRequest(http.MethodPost, baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("HTTP error: " + resp.Status)
	}
	defer resp.Body.Close()

	return nil
}

// postRequest performs a POST request to the endpoint passed as parameter but doesn't send any additional data
func postRequestNoJson(endpoint string) error {
	resp, err := http.Post(baseURL+endpoint, "text/plain", nil)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("HTTP error: " + resp.Status)
	}
	defer resp.Body.Close()

	return nil
}
