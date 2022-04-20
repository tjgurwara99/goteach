package main

import (
	"fmt"
	"log"
	"net/http"
)

func passThroughHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		log.Print("Error creating request: ", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error creating request: ", err)
	}
	ctx := r.Context()
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error making request: ", err)
		w.WriteHeader(500)
	}
	defer resp.Body.Close()
	fmt.Fprint(w, resp.Body)
}

func main() {
	http.HandleFunc("/", passThroughHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
