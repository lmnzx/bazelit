package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type routeServe struct{}

func (rs *routeServe) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	w.Write(fmt.Appendf(nil, "Hello %s", id))
}

func main() {
	port := flag.Int("port", 3000, "Port number to be used for HTTP")
	flag.Parse()

	h := &routeServe{}
	mux := http.NewServeMux()
	mux.Handle("GET /hello/{id}", h)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error running HTTP: %v\n", err)
	}
}
