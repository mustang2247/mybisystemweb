package main

import (
	// "fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!!!!!"))
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
