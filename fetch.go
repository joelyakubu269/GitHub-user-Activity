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
		fmt.Println("Error:", err)
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
