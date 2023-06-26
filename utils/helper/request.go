package helper

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PostRequest(input []byte, url string) ([]byte, error) {
	// Create a new POST client
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(input))
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	var responseBody []byte
	_, err = resp.Body.Read(responseBody)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func GetRequest(url string) ([]byte, error) {
	// Create a new POST client
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
