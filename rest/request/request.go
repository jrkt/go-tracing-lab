package request

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/trace"
)

func GET(url string, span *trace.Span) ([]byte, error) {
	return req("GET", url, span)
}

func POST(url string, span *trace.Span) ([]byte, error) {
	return req("POST", url, span)
}

func req(method, url string, span *trace.Span) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest error: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", method, err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Body error: %s", err)
	}

	resp.Body.Close()

	return b, nil
}
