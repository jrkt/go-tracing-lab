package main

import (
	"flag"
	"log"
	"net/http"

	"cloud.google.com/go/trace"
	"github.com/jrkt/go-tracing-lab/rest/middleware"
	"github.com/jrkt/go-tracing-lab/rest/request"
	"github.com/jrkt/go-tracing-lab/traceclient"
)

func main() {

	port := flag.String("p", "8000", "Port")
	flag.Parse()

	traceClient, err := traceclient.New()
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/SayHello", middleware.TraceRequest(traceClient, sayHello))

	log.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalln(err)
	}
}

func sayHello(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	res, err := request.GET("http://localhost:8001/SayHelloAgain", span)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(res)
}
