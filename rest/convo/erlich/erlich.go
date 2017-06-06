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

	http.HandleFunc("/erlich-1", middleware.TraceRequest(traceClient, erlich1))
	http.HandleFunc("/erlich-2", middleware.TraceRequest(traceClient, erlich2))
	http.HandleFunc("/erlich-3", middleware.TraceRequest(traceClient, erlich3))

	log.Println("Serving on :" + ports.Erlich)
	if err := http.ListenAndServe(":"+ports.Erlich, nil); err != nil {
		log.Fatalln(err)
	}
}

func erlich1(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /erlich-1")

	span.SetLabel("audio", "gilfoyle, whats happening? what?")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/erlich-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /gilfoyle-3")
	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-3", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func erlich2(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /erlich-2")

	span.SetLabel("audio", "go for Erlich..")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/erlich-2.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /jianyang-1")
	_, err = request.POST("http://localhost:"+ports.JianYang+"/jianyang-1", span)
	if err != nil {
		log.Fatalln(err)
	}
}

func erlich3(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /erlich-3")

	span.SetLabel("audio", "I'm gonna let him have this one..")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/erlich-3.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	w.Write([]byte("Conversation finished."))
}
