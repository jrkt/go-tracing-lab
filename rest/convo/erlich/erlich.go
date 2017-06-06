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

	http.HandleFunc("/erlich-1", middleware.TraceRequest(traceClient, erlich1))
	http.HandleFunc("/erlich-2", middleware.TraceRequest(traceClient, erlich2))
	http.HandleFunc("/erlich-3", middleware.TraceRequest(traceClient, erlich3))
	http.HandleFunc("/erlich-4", middleware.TraceRequest(traceClient, erlich4))

	log.Println("Serving on :" + ports.Erlich)
	if err := http.ListenAndServe(":"+ports.Erlich, nil); err != nil {
		log.Fatalln(err)
	}
}

func erlich1(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "gilfoyle, whats happening? what?")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/erlich-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-3", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func erlich2(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "go for Erlich..")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/erlich-2.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.JianYang+"/jianyang-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func erlich3(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "I'm gonna let him have this one..")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/erlich-3.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	err = exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/erlich-4.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.JianYang+"/jianyang-2", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func erlich4(span *trace.Span, w http.ResponseWriter, r *http.Request) {

	span.SetLabel("audio", "Not now JianYang! Go back into your room!")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/erlich-5.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}
}
