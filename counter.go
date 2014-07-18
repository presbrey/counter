package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync/atomic"
)

var (
	addr = flag.String("http", ":8000", "")

	N uint64
)

func init() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	flag.Parse()
}

func main() {
	log.Fatalln((http.Server{
		Addr: *addr,
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			x := atomic.AddUint64(&N, 1)
			fmt.Fprintf(rw, "%d", x)
		}),
	}).ListenAndServe())
}
