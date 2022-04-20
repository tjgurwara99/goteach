package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		fmt.Print("Context canceled")
		if err := ctx.Err(); err != nil {
			log.Print("Context error: ", err)
		}
		w.WriteHeader(500)
	case <-time.After(time.Second * 10):
		fmt.Fprint(w, "handler2: Hello There")
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
