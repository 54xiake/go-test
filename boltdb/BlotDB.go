package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(true)
	defer tx.Rollback()

	if err != nil {
		log.Fatal(err)
	}

	//_, err = tx.CreateBucket([]byte("MyBucket"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	b := tx.Bucket([]byte("MyBucket"))
	err = b.Put([]byte("answer"), []byte("43"))
	if err != nil {
		log.Fatal(err)
	}

	err = b.Put([]byte("answer1"), []byte("431"))
	if err != nil {
		log.Fatal(err)
	}

	v := b.Get([]byte("answer"))
	fmt.Printf("The answer is: %s\n", v)

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
