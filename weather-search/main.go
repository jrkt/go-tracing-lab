package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"cloud.google.com/go/trace"
	"github.com/jonathankentstevens/grpc-tracing-lab/weather-search/weather/client"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	token := flag.String("token", "2b069cd91c2cec1f", "Wunderground API Token")
	zip := flag.Int64("zip", 84651, "Zip Code to retrieve weather conditions")
	flag.Parse()

	// get new weather client
	weatherClient, err := client.New()
	if err != nil {
		log.Fatalf("failed to get weather client: %v", err)
	}

	// establish new trace client
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY_FILE")))
	if err != nil {
		log.Fatalf("failed to establish new trace client: %v", err)
	}

	for i := 0; i < 500; i++ {
		// create root span
		span := tc.NewSpan("/weather/SearchByZip")
		span.SetLabel("zip", strconv.FormatInt(*zip, 10))

		// build span into context
		ctx = trace.NewContext(ctx, span)

		// pass new context into gRPC call to service
		data, err := weatherClient.SearchByZip(ctx, *token, *zip)
		if err != nil {
			log.Fatalf("failed to get weather: %v", err)
		}

		// blocks until traces have been uploaded to gcp
		err = span.FinishWait() // use span.Finish() if your client is a long-running process.
		if err != nil {
			log.Fatalf("error finishing & uploading traces: %v", err)
		}

		fmt.Println("Weather Data:", data)
	}
}
