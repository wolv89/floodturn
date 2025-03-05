package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/wolv89/floodturn/api"
	"github.com/wolv89/floodturn/span"
)

var flagSample int

func main() {

	flag.Parse()

	if flagSample < 0 {
		log.Fatalf("invalid sample: %d", flagSample)
	}

	sp := span.Span{
		Days: make([]span.Day, 0),
	}

	err := sp.Read(flagSample)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	api.LoadRoutes(mux)

	server := &http.Server{
		Handler: mux,
		Addr:    ":35663",
	}

	server.ListenAndServe()

}

func init() {

	flag.IntVar(&flagSample, "sample", 0, "Use sample data file #")
	flag.IntVar(&flagSample, "s", 0, "Use sample data file #")

}
