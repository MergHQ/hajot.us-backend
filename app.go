package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"./resources"
)

func main() {
	println("Starting le web service...")
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	// Registering user resource
	postResource := resources.PostResource{}
	postResource.Register(wsContainer)

	println("Listening..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}