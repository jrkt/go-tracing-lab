package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	port := flag.String("p", "8001", "Port")
	flag.Parse()

	// set up Trace client & endpoint

	log.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalln(err)
	}
}
