package main

import (
	"flag"
	"log"

	"github.com/jonathankentstevens/go-tracing-lab/rest/request"
)

func main() {

	num := flag.Int("n", 1, "Number of requests to make")
	flag.Parse()

	for i := 0; i < *num; i++ {
		res, err := request.GET("http://localhost:8000/SayHello", nil)
		if err != nil {
			log.Fatalln(err)
		}

		println("Response:", string(res))
	}
}
