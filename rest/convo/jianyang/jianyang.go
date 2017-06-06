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

	http.HandleFunc("/jianyang-1", middleware.TraceRequest(traceClient, jianyang1))
	http.HandleFunc("/jianyang-2", middleware.TraceRequest(traceClient, jianyang2))

	log.Println("Serving on :" + ports.JianYang)
	if err := http.ListenAndServe(":"+ports.JianYang, nil); err != nil {
		log.Fatalln(err)
	}
}

func jianyang1(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "Eric a Bachmann.. this is you, as a old man...")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/jianyang-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Erlich+"/erlich-3", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func jianyang2(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "Eric Bachmann..this is your mom..")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/jianyang-2.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Erlich+"/erlich-4", span)
	if err != nil {
		log.Fatalln(err)
	}
}
