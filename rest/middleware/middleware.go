package middleware

import (
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/trace"
)

type traceHandlerFunc func(*trace.Span, http.ResponseWriter, *http.Request)

func TraceRequest(tc *trace.Client, fn traceHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		fmt.Println("received request:", r.RequestURI)

		span := tc.SpanFromRequest(r)
		defer span.Finish()

		time.Sleep(5 * time.Millisecond)

		fn(span, w, r)
	})
}
