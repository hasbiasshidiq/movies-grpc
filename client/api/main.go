package main

import (
	"omdb-client/api/handler"
	"omdb-client/api/middleware"
	"omdb-client/config"
	grpcdriver "omdb-client/infrastructure/grpc-driver"
	movies "omdb-client/usecase/movies"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {

	config.LoadEnv()

	r := mux.NewRouter().StrictSlash(true)

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	omdbGRPC := grpcdriver.NewOmdbGRPC()

	moviesService := movies.NewService(omdbGRPC)

	handler.MakeMoviesHandlers(r, *n, moviesService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + config.API_PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	log.Printf("\nServer starting on port %s", config.API_PORT)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
