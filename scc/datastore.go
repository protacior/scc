package scc

import (
	"appengine"
	"appengine/datastore"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Verified  bool
}

// userKey returns the key used for a specific user email address
func userKey(c appengine.Context, user string) *datastore.Key {
	return datastore.NewKey(c, "User", user, 0, nil)
}

func getUser(c appengine.Context, email string) (usr *User, err error) {
	key := datastore.NewKey(c, "User", email, 0, nil)
	err = datastore.Get(c, key, &usr)
	return usr, err
}

func storeUser(c appengine.Context, usr User) error {
	key := datastore.NewKey(c, "User", usr.Email, 0, nil)
	_, err := datastore.Put(c, key, &usr)
	return err
}
