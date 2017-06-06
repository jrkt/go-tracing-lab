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

	http.HandleFunc("/dinesh-1", middleware.TraceRequest(traceClient, dinesh1))

	log.Println("Serving on :" + ports.Dinesh)
	if err := http.ListenAndServe(":"+ports.Dinesh, nil); err != nil {
		log.Fatalln(err)
	}
}

func dinesh1(span *trace.Span, w http.ResponseWriter, r *http.Request) {
	span.SetLabel("audio", "you said delete..")

	err := exec.Command("cvlc", "--play-and-exit", os.Getenv("GOPATH")+"/src/github.com/jonathankentstevens/go-tracing-lab/audio/dinesh-1.mp3").Run()
	if err != nil {
		log.Fatalln("failed to play audio:", err)
	}

	_, err = request.POST("http://localhost:"+ports.Gilfoyle+"/gilfoyle-2", span)
	if err != nil {
		log.Fatalln(err)
	}
}
