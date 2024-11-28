package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = 8081

type Config struct{}

func main() {
	app := Config{}

	log.Println("Starting server on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
