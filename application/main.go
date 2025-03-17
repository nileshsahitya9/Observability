package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server is starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Root endpoint hit")
		fmt.Fprintf(w, "Welcome to the Go Server")
	})

	http.HandleFunc("/service-a", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Service A endpoint hit")
		fmt.Fprintf(w, "Service A is running")
	})

	http.HandleFunc("/service-b", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Service B endpoint hit")
		fmt.Fprintf(w, "Service B is running")
	})

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
