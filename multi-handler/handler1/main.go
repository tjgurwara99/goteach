package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	if err != nil {
		log.Printf("handler1 err: %v", err)
		w.WriteHeader(500)
		return
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Printf("handler1 canceled: %v", err)
			w.WriteHeader(500)
			return
		}
		log.Printf("handler1 err: %v", err)
		w.WriteHeader(500)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("handler1 err: %v", err)
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w, "handler1: %s", body)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
