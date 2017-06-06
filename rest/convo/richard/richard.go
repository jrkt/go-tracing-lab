package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"cloud.google.com/go/trace"
	"github.com/jonathankentstevens/go-tracing-lab/rest/convo/ports"
	"github.com/jonathankentstevens/go-tracing-lab/rest/middleware"
	"github.com/jonathankentstevens/go-tracing-lab/rest/request"
	"github.com/jonathankentstevens/go-tracing-lab/traceclient"
)

func main() {

	traceClient, err := traceclient.New()
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/startconvo", middleware.TraceRequest(traceClient, startconvo))
	http.HandleFunc("/richard-2", middleware.TraceRequest(traceClient, richard2))
	http.HandleFunc("/richard-3", middleware.TraceRequest(traceClient, richard3))
	http.HandleFunc("/richard-4", middleware.TraceRequest(traceClient, richard4))

	log.Println("Serving on :" + ports.Richard)
	if err := http.ListenAndServe(":"+ports.Richard, nil); err != nil {
		log.Fatalln(err)
	}
}

func startconvo(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "did you delete it?")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/richard-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func richard2(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "oops, what does that mean, oops?")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/richard-2.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Dinesh+"/dinesh-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func richard3(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "what?")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/richard-3.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Erlich+"/erlich-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func richard4(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "so what does that mean? did you delete Pied Piper or not?")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/richard-4.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-4", span)
	if err != nil {
		log.Fatalln(err)
	}
}
