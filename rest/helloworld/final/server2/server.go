package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/trace"
	"github.com/jonathankentstevens/go-tracing-lab/rest/middleware"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	port := flag.String("p", "8001", "Port")
	flag.Parse()

	traceClient, err := trace.NewClient(context.Background(), os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		log.Fatalf("\nError creating trace client: %s", err)
	}

	http.HandleFunc("/SayHelloAgain", middleware.TraceRequest(traceClient, sayHelloAgain))

	log.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalln(err)
	}
}

func sayHelloAgain(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /SayHelloAgain")

	time.Sleep(5 * time.Millisecond)

	w.Write([]byte("Hello Again"))
}
