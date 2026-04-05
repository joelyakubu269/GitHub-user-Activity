package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

func fetchApi(username string) ([]Event, error) {
	var usernam string
	flag.StringVar(&usernam, "user", "", "github username to lookup info on.")
	flag.Parse()

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {

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
