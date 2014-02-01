package scc

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Verified  bool
}

type Request struct {
	Name          string
	Email         string
	StartDate     Time
	EndDate       Time
	LowPrice      int
	HighPrice     int
	City          string
	State         string
	Gender        string
	GenderDesired string
	Desired       string
	Am            string
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

// same for "Room Request" but also with sorting (chunks?)
// Queries : Google App Engine Go
