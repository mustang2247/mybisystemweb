package main

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/mustang2247/mybisystemweb/httpservice"
)

func main() {
	log.Print("Starting server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!!!!!"))
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
