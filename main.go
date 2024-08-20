package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infra-monkey/go-google-calendar-notify/gcal"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := gcal.NewServer()

	r := gin.Default()

	gcal.RegisterHandlers(r, server)

	r.Handle("GET", "/", server.GetHome)
	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())

	// Create a fake authenticator. This allows us to issue tokens, and also
	// implements a validator to check their validity.
	// authenticator, err := api.NewJwtAuthenticator()
	// if err != nil {
	// 	log.Fatalln("error creating authenticator:", err)
	// }

	// mw, err := api.CreateMiddleware(authenticator)
	// if err != nil {
	// 	log.Fatalln("error creating middleware:", err)
	// }

	// fs := http.FileServer(http.Dir("public"))
	// http.Handle("/gcal/", fs)

	// h = mw(h)
}
