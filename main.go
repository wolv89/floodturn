package main

import (
	"flag"
	"log"

	"github.com/wolv89/floodturn/api"
	"github.com/wolv89/floodturn/span"
)

var flagSample int

func main() {

	api.Run()

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

	sp.Render()

}

func init() {

	flag.IntVar(&flagSample, "sample", 0, "Use sample data file #")
	flag.IntVar(&flagSample, "s", 0, "Use sample data file #")

}
