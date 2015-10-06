package main

import (
	"flag"
	"net/http"
	"router"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// VERSION - service version. Value is replaced during the build process with
// build date and last git commit short hash. For example: 201510061226@b3c3271
var VERSION = ""

func init() {
	flag.Parse()
	router.Config.Prepare()
}

func main() {
	glog.Infof("Listening on 0.0.0.0:8080")
	r := mux.NewRouter()
	r.HandleFunc("/", router.PageIndex)
	r.HandleFunc("/router-stats", router.PageStats)

	routedHandlerMain := http.HandlerFunc(router.PageRouted)
	routedHandlerChain := alice.New(router.CheckRequestDigest).Then(routedHandlerMain)
	r.Handle("/{digest}/{url}", routedHandlerChain)

	http.ListenAndServe(":8080", r)
}
