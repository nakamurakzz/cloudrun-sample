package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Sample HTTP Server
	log.Println("Starting HTTP Server")
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	w.Write([]byte("Hello, " + name))
}
