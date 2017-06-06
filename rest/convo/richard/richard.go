package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"cloud.google.com/go/trace"
	"github.com/jonathankentstevens/go-tracing-lab/rest/convo/ports"
	"github.com/jonathankentstevens/go-tracing-lab/rest/middleware"
	"github.com/jonathankentstevens/go-tracing-lab/rest/request"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	traceClient, err := trace.NewClient(context.Background(), os.Getenv("GCP_PROJECT"), option.WithServiceAccountFile(os.Getenv("GCP_KEY")))
	if err != nil {
		log.Fatalf("Error creating trace client: %s", err)
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
	fmt.Println("received request: /startconvo")

	span.SetLabel("audio", "did you delete it?")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/richard-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /gilfoyle-1")
	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func richard2(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /richard-2")

	span.SetLabel("audio", "oops, what does that mean, oops?")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/richard-2.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Dinesh+"/dinesh-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func richard3(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /richard-3")

	span.SetLabel("audio", "what?")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/richard-3.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Erlich+"/erlich-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func richard4(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /richard-4")

	span.SetLabel("audio", "so what does that mean? did you delete Pied Piper or not?")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/richard-4.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-4", span)
	if err != nil {
		log.Fatalln(err)
	}
}
