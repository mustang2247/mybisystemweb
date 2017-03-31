package main

import (
	"log"

	// "./routers/httpservice"
	"github.com/mustang2247/mybisystemweb/routers/httpservice"
)

func main() {
	log.Print("Starting server")

	httpservice.InitHttp()
}
