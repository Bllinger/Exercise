package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "into handler1\n")
	mu.Lock()
	//count = count+1
	fmt.Fprintf(w, "conut++\n")
	mu.Unlock()
	fmt.Fprintf(w, "Into URL.PATH=%v\n", r.URL.Path)
	fmt.Fprintf(w, "%s\t%s\n", r.Method, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "into counter\n")
	mu.Lock()
	fmt.Fprintf(w, "Into Counter!\tcount = %v\n", count)
	mu.Unlock()
}
