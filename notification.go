package main

import (
	"context"
	"log"
	"time"

	"github.com/godbus/dbus/v5"
)

func notify(ctx context.Context, notification Notification) {
	duration := time.Duration(notification.Interval) * time.Second
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sendNotification(notification.Header, notification.Message)
		case <-ctx.Done():
			return
		}
	}
}

func sendNotification(title, message string) {
	playMP3("assets/sounds/chime.mp3")

	conn, err := dbus.SessionBus()
	if err != nil {
		log.Println("Failed to connect to session bus:", err)
		return
	}

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0,
		"PauseMate",
		uint32(0),
		"",
		title,
		message,
		[]string{},
		map[string]dbus.Variant{},
		int32(5000),
	)
	if call.Err != nil {
		panic(call.Err)
	}
}
