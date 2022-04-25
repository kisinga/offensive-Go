package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type logger struct {
	inner http.Handler
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}
func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("start")
	l.inner.ServeHTTP(w, req)
	time.Sleep(time.Second * 3)
	log.Println("Finish")
}

func main() {
	f := http.HandlerFunc(hello)
	lo := logger{
		inner: f,
	}
	http.ListenAndServe(":8080", &lo)
}
