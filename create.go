package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Reset = "\033[0m"
var Red = "\033[31m"

func CreateNotification(path bool) {
	fmt.Println("Creating...")

	message := getUserInput("Enter a message for yourself.")
	header := getUserInput("Enter a header for the message.")

	var interval int
	for {
		intervalInput := getUserInput("Enter an interval for the message in minutes.")
		var err error
		interval, err = strconv.Atoi(intervalInput)
		if err != nil {
			fmt.Println(Red + "Invalid interval. Please enter a valid number.\n" + Reset)
		} else {
			break
		}
	}

	fmt.Println("Message: ", message)
	fmt.Println("Header: ", header)
	fmt.Println("Interval: ", interval)

	notification := Notification{
		Message:  message,
		Header:   header,
		Interval: interval,
		Path:     "assets/images/default.png",
	}
	fmt.Println(notification)

	WriteNotificationJSON("config/notifications.json", notification)
}

func getUserInput(message string) string {
	fmt.Println(message)
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}
