package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func someFunctionFails() error {
	return errors.New("SomeError")
}

func someLongFunction(cancel context.CancelFunc) string {
	err := someFunctionFails()
	if err != nil {
		cancel()
	}
	<-time.After(5 * time.Second)
	return "Finished"
}

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, cancel := context.WithCancel(ctx)
		logger.Print("processing request")
		message := make(chan string, 1)

		go func() {
			message <- someLongFunction(cancel)
		}()
		select {
		case m := <-message:
			logger.Print("Message received")
			fmt.Fprintf(w, m)
		case <-ctx.Done():
			logger.Print("request cancelled")
			fmt.Fprintf(w, "Message cancelled")
		}
	}))
}
