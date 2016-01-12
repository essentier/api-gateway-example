package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/codegangsta/negroni"
	"github.com/essentier/spickspan"
	"github.com/gorilla/mux"
)

func main() {
	router := initRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Printf("Listening on port 8087")
	log.Fatal(http.ListenAndServe(":8087", n))
}

func initRoutes() *mux.Router {
	provider, err := spickspan.GetDefaultServiceProvider()
	if err != nil {
		log.Fatalf("Could not resolve spickspan provider. The error is %v", err)
		return nil
	}

	todoService, err := spickspan.GetHttpService(provider, "todo-rest", "/todos")
	if err != nil {
		log.Fatalf("Could not get the todo service. The error is %v", err)
		return nil
	}

	helloService, err := spickspan.GetHttpService(provider, "hello-rest", "/hello")
	if err != nil {
		log.Fatalf("Could not get the hello service. The error is %v", err)
		return nil
	}

	router := mux.NewRouter()
	setReverseProxyRoutes(router, todoService.GetUrl(), "/todo-rest/")
	setReverseProxyRoutes(router, helloService.GetUrl(), "/hello-rest/")
	return router
}

func setReverseProxyRoutes(router *mux.Router, targetUrl string, prefix string, handlers ...negroni.Handler) {
	target, err := url.Parse(targetUrl)
	if err != nil {
		log.Fatal(err)
	}

	targetHandler := negroni.Wrap(http.StripPrefix(prefix, httputil.NewSingleHostReverseProxy(target)))
	allHandlers := append(handlers, targetHandler)
	router.Handle(prefix+"{rest:.*}", negroni.New(allHandlers...))
}
