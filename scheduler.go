package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Concurrrent System that will count start our notifications on their loop.
Parameters:

The function sets up a context with cancellation and listens for system
interrupt signals (SIGINT, SIGTERM) to gracefully shut down the notification
loops. Each notification is processed in a separate goroutine, and the
function waits for a termination signal to cancel the context and exit the
program.
*/
func ConcurentSystem(notifications []Notification) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for _, notification := range notifications {
		go notify(ctx, notification)
	}

	fmt.Println("Press Ctrl+C to exit...")

	<-sigChan

	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Program exiting.")
}
