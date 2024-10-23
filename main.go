package main

import (
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

		ConcurentSystem(notifications)
	}
}

// ᓚᘏᗢ
