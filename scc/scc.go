package scc

import (
	"appengine"
	"appengine/user"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/signin/", SignInHandler)
	http.HandleFunc("/signup/", SignUpHandler)
	http.HandleFunc("/search/", SearchHandler)
	http.HandleFunc("/url/", VerHandler)
}

func VerHandler(w http.ResponseWriter, r *http.Request) {
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, %v!", u)
}
