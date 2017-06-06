package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	port := flag.String("p", "8000", "Port")
	flag.Parse()

	// initialize Trace client & http endpoint

	log.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalln(err)
	}
}
