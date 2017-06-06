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

	http.HandleFunc("/gilfoyle-1", middleware.TraceRequest(traceClient, gilfoyle1))
	http.HandleFunc("/gilfoyle-2", middleware.TraceRequest(traceClient, gilfoyle2))
	http.HandleFunc("/gilfoyle-3", middleware.TraceRequest(traceClient, gilfoyle3))
	http.HandleFunc("/gilfoyle-4", middleware.TraceRequest(traceClient, gilfoyle4))

	log.Println("Serving on :" + ports.Gilfoyle)
	if err := http.ListenAndServe(":"+ports.Gilfoyle, nil); err != nil {
		log.Fatalln(err)
	}
}

func gilfoyle1(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /gilfoyle-1")

	span.SetLabel("audio", "oops")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/gilfoyle-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /richard-2")
	_, err = request.POST("http://localhost:"+ports.Richard+"/richard-2", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func gilfoyle2(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /gilfoyle-2")

	span.SetLabel("audio", "oh wait, hold on..")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/gilfoyle-2.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /richard-3")
	_, err = request.POST("http://localhost:"+ports.Richard+"/richard-3", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func gilfoyle3(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /gilfoyle-3")

	span.SetLabel("audio", "thats weird, kernel panic, the whole system just shut down")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/gilfoyle-3.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /richard-4")
	_, err = request.POST("http://localhost:"+ports.Richard+"/richard-4", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func gilfoyle4(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /gilfoyle-4")

	span.SetLabel("audio", "nope, we definitely did not. thanks to Dinesh's garbage code...")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/gilfoyle-4.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /erlich-2")
	_, err = request.POST("http://localhost:"+ports.Erlich+"/erlich-2", span)
	if err != nil {
		log.Fatalln(err)
	}
}
