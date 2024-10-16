package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "-c":
			path := false
			CreateNotification(path)
		}
	} else {
		notifications := ReadNotificationJSON("config/notifications.json")
		fmt.Println(notifications)
		// Start the notifications
		ConcurentSystem(notifications)
	}
}

// ᓚᘏᗢ

// Remember treyson,
// Keep main.go clean and simple.
// Compartmentalize your code into different files.
// Keep users in mind when designing your CLI.

// ᓚᘏᗢ
