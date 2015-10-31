package main

import (
	"log"
	"net/http"
)

func main() {

	db, err := NewOpen("mysql", "jake:password@/pfm")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	router := NewRouter(db)

	http.ListenAndServe(":7012", router)
}
