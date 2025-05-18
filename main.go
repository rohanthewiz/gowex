package main

import (
	"github.com/rohanthewiz/rweb"
	"log"
)

func main() {
	// Start the server
	s := rweb.NewServer(
		rweb.ServerOptions{
			Address: "localhost:8000",
			Verbose: true,
		})

	// Middleware
	s.Use(rweb.RequestInfo)

	htmlHandler(s)
	exeHandler(s)
	fmtHandler(s)

	log.Fatal(s.Run())
}
