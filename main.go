package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Starting server")


	//httpservice.InitHttpService()



	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!!!!!"))
	})

	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello aa"))
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
