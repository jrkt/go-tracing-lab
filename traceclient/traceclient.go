package traceclient

import (
	"fmt"
	"os"

	"cloud.google.com/go/trace"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func New() (*trace.Client, error) {
	client, err := trace.NewClient(context.Background(), os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		return nil, fmt.Errorf("Error creating trace client: %s", err)
	}

	p, err := trace.NewLimitedSampler(1, 1000)
	if err != nil {
		return nil, fmt.Errorf("Error setting trace sampler: %s", err)
	}
	client.SetSamplingPolicy(p)

	return client, nil
}
