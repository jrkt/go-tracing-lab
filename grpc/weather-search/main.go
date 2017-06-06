package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/trace"
	"github.com/jonathankentstevens/go-tracing-lab/grpc/weather-search/weather/client"
	"github.com/jonathankentstevens/go-tracing-lab/traceclient"
	"golang.org/x/net/context"
)

func main() {

	token := flag.String("token", "2b069cd91c2cec1f", "Wunderground API Token")
	zip := flag.Int64("zip", 84651, "Zip Code to retrieve weather conditions")
	num := flag.Int("n", 1, "Number of times to run request")
	flag.Parse()

	// get new weather client
	weatherClient, err := client.New()
	if err != nil {
		log.Fatalf("failed to get weather client: %v", err)
	}

	// establish new trace client
	ctx := context.Background()
	tc, err := traceclient.New()
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < *num; i++ {
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
