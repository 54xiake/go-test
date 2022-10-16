package main

import (
	"fmt"
	"github.com/rhinoman/couchdb-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

type TestDocument struct {
	Title string
	Note  string
}

func main() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	if err != nil {
		fmt.Println(err)
	}
	auth := couchdb.BasicAuth{Username: "admin", Password: "adminpw"}
	db := conn.SelectDB("my_database", &auth)

	theDoc := TestDocument{
		Title: "My Document",
		Note:  "This is a note",
	}
	for i := 0; i < 10; i++ {
		//theId := uuid.Must(uuid.NewV4())
		theId := uuid.NewV4()

		//The third argument here would be a revision, if you were updating an existing document
		rev, err := db.Save(theDoc, theId.String(), "")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(rev)
	}
	//If all is well, rev should contain the revision of the newly created
	//or updated Document
}
