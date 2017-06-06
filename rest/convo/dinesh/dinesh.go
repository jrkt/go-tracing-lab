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

	http.HandleFunc("/dinesh-1", middleware.TraceRequest(traceClient, dinesh1))

	log.Println("Serving on :" + ports.Dinesh)
	if err := http.ListenAndServe(":"+ports.Dinesh, nil); err != nil {
		log.Fatalln(err)
	}
}

func dinesh1(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: /dinesh-1")

	span.SetLabel("audio", "you said delete..")

	err := exec.Command("cvlc", "--play-and-exit", "/home/jstevens/Presentation/dinesh-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	fmt.Println("making request: /gilfoyle-2")
	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-2", span)
	if err != nil {
		log.Fatalln(err)
	}
}
