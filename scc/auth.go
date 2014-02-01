func auth(w http.ResponseWriter, r *http.Request) {
	// show "signing page"
	// must get email and password correct
	c := appengine.NewContext(r)
	usr, err := getUser(c, "test@email.com")
	if err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// check if email and authentication link match
	url := r.URL
	code := url.Query()
	if code == usr.Auth {
		usr.Verified = true
		usr.Auth = ""
	}
	err = storeUser(c, usr)
	if err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
