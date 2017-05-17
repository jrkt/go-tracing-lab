package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var traceClient *trace.Client

func main() {

	port := flag.String("p", "8080", "Port")
	flag.Parse()

	var err error
	traceClient, err = trace.NewClient(context.Background(), os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		log.Fatalf("\nError creating trace client: %s", err)
	}

	http.HandleFunc("/SayHello", traceRequest(sayHello))
	http.HandleFunc("/SayHelloAgain", traceRequest(sayHelloAgain))

	log.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalln(err)
	}
}

func traceRequest(fn http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		span := traceClient.SpanFromRequest(r)
		defer span.Finish()

		fn(w, r)
	})
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func sayHelloAgain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Again"))
}
