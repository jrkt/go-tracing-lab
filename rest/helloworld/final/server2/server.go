package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/trace"
	"github.com/jonathankentstevens/go-tracing-lab/rest/middleware"
	"github.com/jonathankentstevens/go-tracing-lab/traceclient"
)

func main() {

	port := flag.String("p", "8001", "Port")
	flag.Parse()

	traceClient, err := traceclient.New()
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/SayHelloAgain", middleware.TraceRequest(traceClient, sayHelloAgain))

	log.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalln(err)
	}
}

func sayHelloAgain(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Millisecond)

	w.Write([]byte("Hello Again"))
}
