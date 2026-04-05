package main

import (
	"flag"
	"fmt"
)

func main() {
	var username string
	flag.StringVar(&username, "user", "", "github  username to lookup info on.")
	flag.BoolVar(&logEnabled, "loud", false, "set to enable verbose logging")
	flag.Parse()
	if flag.NArg() > 0 {
		fmt.Println("no arguements expected")
		return
	}
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

var logEnabled bool

func logstuff(args ...any) {
	if logEnabled {
		fmt.Println(args...)
	}
}
