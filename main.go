package main

import (
	"log"
	"net/http"
)

func main() {

	go AutoResolveHost()

	// start web server
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
