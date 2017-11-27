package main

import (
	"log"

	"github.com/jrkt/go-tracing-lab/rest/convo/ports"
	"github.com/jrkt/go-tracing-lab/rest/request"
)

func main() {

	_, err := request.POST("http://localhost:"+ports.Richard+"/startconvo", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
