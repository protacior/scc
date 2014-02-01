package scc

import (
	"appengine"
	"net/http"
)

func auth(w http.ResponseWriter, r *http.Request) {
	// show "signing page"
	// must get email and password correct
	c := appengine.NewContext(r)
	usr, err := getUser(c, "test@email.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// check if email and authentication link match
	url := r.URL
	code := url.RawQuery
	if code == usr.Auth {
		usr.Verified = true
		usr.Auth = ""
	}
	// update the user's information online
	err = storeUser(c, usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
