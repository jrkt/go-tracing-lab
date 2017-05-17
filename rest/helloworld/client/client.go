package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var traceClient *trace.Client

func main() {

	var err error
	traceClient, err = trace.NewClient(context.Background(), os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		log.Fatalf("\nError creating trace client: %s", err)
	}

	span := traceClient.NewSpan("SayHello")
	defer span.FinishWait()

	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/SayHello", nil)
	if err != nil {
		log.Fatalf("\nNewRequest (SayHello) error: %s", err)
	}
	req.Header.Add("X-Cloud-Trace-Context", fmt.Sprintf("%s/%d;o=1", span.TraceID(), 0))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("\nGET (SayHello) error: %s", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("\nBody (SayHello) error: %s", err)
	}

	println("Response:", string(b))
}
