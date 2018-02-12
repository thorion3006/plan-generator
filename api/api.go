package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thorion3006/plan-generator/api/middleware"

	"github.com/thorion3006/plan-generator/api/controller"
)

// Creates a http Server on port 8080.
func startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: &middleware.TimeoutMiddleware{Next: new(middleware.GzipMiddleware)}}

	controller.Setup()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	fmt.Println("Web server is listening on port 8080...")
	// returning reference so caller can call Shutdown()
	return srv
}

// Starts the web server for the plangenerator.
func Start() {
	srv := startHttpServer()
	fmt.Println("Press Enter to shutdown the server...")
	var shutdown string
	fmt.Scanln(&shutdown)
	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}
	fmt.Println("Thank you for using plan-generator!")
}
