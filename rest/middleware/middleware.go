package middleware

import (
	"net/http"
	"time"

	"cloud.google.com/go/trace"
)

type traceHandlerFunc func(*trace.Span, http.ResponseWriter, *http.Request)

func TraceRequest(tc *trace.Client, fn traceHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		span := tc.SpanFromRequest(r)
		defer span.FinishWait()

		time.Sleep(5 * time.Millisecond)

		fn(span, w, r)
	})
}
