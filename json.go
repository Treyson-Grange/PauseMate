package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Notification struct {
	Message  string `json:"message"`
	Path     string `json:"path"`
	Header   string `json:"header"`
	Interval int    `json:"interval"`
}

func ReadNotificationJSON(filePath string) []Notification {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while opening", filePath, " : "+err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data = []Notification{}
	if err := decoder.Decode(&data); err != nil {
		log.Fatal("Error while decoding JSON: ", err)
	}
	return data
}

func WriteNotificationJSON(filePath string, notif Notification) bool {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while opening", filePath, " : "+err.Error())
		return false
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("failed to read file: %w", err)
		return false
	}

	var notifications []Notification
	if err := json.Unmarshal(data, &notifications); err != nil {
		log.Fatal("failed to unmarshal JSON: %w", err)
		return false
	}

	// Append the new notification
	notifications = append(notifications, notif)
	updatedData, err := json.MarshalIndent(notifications, "", "  ")
	if err != nil {
		log.Fatal("failed to marshal JSON: %w", err)
		return false
	}

	// Write it again
	if err := os.WriteFile(filePath, updatedData, 0644); err != nil {
		log.Fatal("failed to write to file: %w", err)
		return false
	}

	return true

}
