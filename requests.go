package cidergo

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
)

const baseURL = "http://0.0.0.0:10767/api/v1/playback/"

var apiToken = ""

func SetToken(token string) {
	apiToken = token
}

// statusRequest performs a GET request to the endpoint passed as a parameter and returns the HTML status as int
func statusRequest(endpoint string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("apptoken", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return resp.StatusCode, errors.New("error: Token required but not set. Please use SetToken or disable API tokens in Cider. More information can be found in the docs")
	} else if resp.StatusCode != http.StatusOK {
		return 0, errors.New("HTTP error: " + resp.Status)
	}

	return resp.StatusCode, nil
}

// jsonRequest performs a GET request to the endpoint passed as a parameter and returns the JSON received from that request
func jsonRequest(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("apptoken", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return nil, errors.New("error: Token required but not set. Please use SetToken or disable API tokens in Cider. More information can be found in the docs")
	} else if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP error: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return body, nil
}

// postRequest performs a POST request to the endpoint passed as a parameter and sends the jsonData (in bytes) passed as a parameter
func postRequest(endpoint string, jsonData []byte) error {
	req, err := http.NewRequest(http.MethodPost, baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apptoken", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return errors.New("error: Token required but not set. Please use SetToken or disable API tokens in Cider. More information can be found in the docs")
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP error: " + resp.Status)
	}

	return nil
}

// postRequestNoJson performs a POST request to the endpoint passed as a parameter but doesn't send any additional data
func postRequestNoJson(endpoint string) error {
	req, err := http.NewRequest(http.MethodPost, baseURL+endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("apptoken", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return errors.New("error: Token required but not set. Please use SetToken or disable API tokens in Cider. More information can be found in the docs")
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP error: " + resp.Status)
	}

	return nil
}
