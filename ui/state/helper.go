package state

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Get retrieve url and return response. If Get encounters an error, it will log and return the error.
func Get(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error Get() in url:%v, err:%v", url, err)
		return nil, err
	}
	return res, nil
}

// Post marshals the payload and send an HTTP POST request with the payload and return the response.
// If Post encounters an error, it will log and return the error.
func Post(url string, contentType string, payload interface{}) (*http.Response, error) {
	jsonValue, _ := json.Marshal(payload)
	res, err := http.Post(url, contentType, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("Error Post() in url:%v err:%v", url, err)
		return res, err
	}
	return res, nil
}

// Patch marshals the payload and send an HTTP Patch request with the payload and return the response.
// If Patch encounters an error, it will log and return the error.
func Patch(url string, payload interface{}) (*http.Response, error) {
	client := http.DefaultClient
	jsonValue, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("Error Patch() in url:%v err:%v", url, err)
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error Patch() in url:%v err:%v", url, err)
		return nil, err
	}
	return resp, nil
}

//Decoder decodes response body to s
func Decoder(res *http.Response, s interface{}) (err error) {
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&s)

	if err != nil {
		log.Printf("Error JSON decoding: %v", err)
		return err
	}
	return nil
}
