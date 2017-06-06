package request

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/trace"
)

func GET(url string, span *trace.Span) ([]byte, error) {
	return do("GET", url, span)
}

func POST(url string, span *trace.Span) ([]byte, error) {
	return do("POST", url, span)
}

func do(method, url string, span *trace.Span) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest error: %s", err)
	}
	fmt.Println("making request:", req.URL.RequestURI())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", method, err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Body error: %s", err)
	}

	return b, nil
}
