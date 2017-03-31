package httpservice

import (
	"log"
	"net/http"

	//"gopkg.in/macaron.v1"
	//"github.com/go-macaron/macaron"
	"macaron"
)

func InitHttpService() {
	m := macaron.Classic()
	m.Get("/", myHandler)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func myHandler(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}
