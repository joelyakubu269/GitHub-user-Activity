package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run . <github-username>")
		return
	}
	username := os.Args[1]
	event, err := fetchApi(username)
	if err != nil {
		fmt.Println("Error fetching events:", err)
		return
	}
	if len(event) == 0 {
		fmt.Println("There is no event to be fetched")
		return
	}
	printEvent(event)
}
