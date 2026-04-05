package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchApi(username string) ([]Event, error) {

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {

		return nil, err
	}
	logstuff(resp.StatusCode)
	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return events, nil
}
