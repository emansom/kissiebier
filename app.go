package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/timshannon/badgerhold"
)

func main() {
	options := badgerhold.DefaultOptions
	options.Dir = "data"
	options.ValueDir = "data"

	conn, err := badgerhold.Open(options)

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	defer conn.Close()

	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Use(DatabaseConnCtx(conn))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe(":3000", mux)
}
