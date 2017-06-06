package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/trace"
	"github.com/jonathankentstevens/go-tracing-lab/rest/middleware"
	"github.com/jonathankentstevens/go-tracing-lab/rest/request"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	port := flag.String("p", "8000", "Port")
	flag.Parse()

	traceClient, err := trace.NewClient(context.Background(), os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		log.Fatalf("\nError creating trace client: %s", err)
	}

	http.HandleFunc("/SayHello", middleware.TraceRequest(traceClient, sayHello))

	log.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalln(err)
	}
}

func sayHello(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /SayHello")

	fmt.Println("making request: /SayHelloAgain")
	res, err := request.GET("http://localhost:8001/SayHelloAgain", span)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(res)
}
