package main

import (
	"github.com/HouzuoGuo/tiedot/db"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	dir := "tiedotDB"
	tdDB, err := db.OpenDB(dir)
	defer tdDB.Close()
	if err != nil {
		panic(err)
	}
	if err := tdDB.Create("mobclients", 2); err != nil {

		panic(err)
	
	} else {
	
		mobclients := tdDB.Use("mobclients")	
		if err := mobclients.Index([]string{"ClPhonenum"}); err != nil {
			panic(err)
		}
	
	}
	
}
