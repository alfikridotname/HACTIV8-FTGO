package main

import (
	"fmt"
	"net/http"

	"net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/loop", loop)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/heap", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	fmt.Println("listening..")
	fmt.Println(http.ListenAndServe(":4560", r))
}

func loop() {
	for i := 0; i < 100000; i++ {
		fmt.Println("test")
	}
}
