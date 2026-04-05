package main

import (
	"flag"
	"fmt"
)

func main() {
	var username string
	var maxEvents int
	flag.StringVar(&username, "user", "", "github  username to lookup info on.")
	flag.BoolVar(&logEnabled, "useloud", false, "set to enable verbose logging")
	flag.IntVar(&maxEvents, "maxevents", 5, "limit maximum nummber of events printed")
	flag.Parse()
	if flag.NArg() > 0 {
		fmt.Println("no arguements expected")
		return
	}
	events, err := fetchApi(username)
	if err != nil {
		fmt.Println("Error fetching events:", err, "for user", username)
		return
	}
	if len(events) > maxEvents {
		events = events[len(events)-5:]
	}
	printEvent(events)
}

var logEnabled bool

func logstuff(args ...any) {
	if logEnabled {
		fmt.Println(args...)
	}
}
