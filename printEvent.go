package main

import "fmt"

func printEvent(event []Event) {
	start := 0

	if len(event) > 5 {
		start = len(event) - 5
	}

	for _, r := range event[start:] {
		fmt.Printf("%s --> %s\n", r.Type, r.Repo.Name)
		commits := r.Payload.Commits
		if r.Type == "PushEvent" && len(commits) > 0 {
			start := 0
			if len(commits) > 5 {
				start = len(commits) - 5
			}
			for _, c := range commits[start:] {

				fmt.Printf("commit: %s\n", c.Message)
			}
		}

	}
}
