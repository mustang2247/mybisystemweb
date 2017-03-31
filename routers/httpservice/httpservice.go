package httpservice

import (
	"log"
	"net/http"
	// "macaron"
)

func InitHttpService() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!!!!!"))
	})

	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello aa"))
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
	// m := macaron.Classic()
	// m.Get("/", myHandler)

	// log.Println("Server is running...")
	// log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

// func myHandler(ctx *macaron.Context) string {
// 	return "the request path is: " + ctx.Req.RequestURI
// }
