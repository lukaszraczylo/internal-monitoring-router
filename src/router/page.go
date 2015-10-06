package router

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// CheckRequestDigest - it's a middleware executed before every query
// to verify that other side has provided correct digest hash.
func CheckRequestDigest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		v := CompareDigest(vars["url"], vars["digest"])
		if !v {
			glog.Errorf("INVALID DIGEST from %s for %s", r.RemoteAddr, vars["url"])
			// Return 404 if invalid digest has been detected
			http.Error(w, http.StatusText(404), 404)
		} else {
			// If digest is exactly as expected - route to the specified URL
			next.ServeHTTP(w, r)
		}
	})
}

// PageIndex - throws regular index.
func PageIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "MoJDS internal monitoring router.")
}

// PageRouted - throws routeg page content.
func PageRouted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	decodedURL, err := base64.StdEncoding.DecodeString(vars["url"])
	if err != nil {
		glog.Errorf("Failed to decode url parameter: %s", vars["url"])
		fmt.Fprintf(w, "Invalid URL specified. Is it base64 encoded one?")
	} else {
		resp, err := http.Get(fmt.Sprintf("%s", decodedURL))
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("%s", err))
			http.Error(w, http.StatusText(500), 500)
		} else {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Fprintf(w, fmt.Sprintf("%s", body))
		}
	}
}

// PageStats - displays routing stats.
func PageStats(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Stats")
}
