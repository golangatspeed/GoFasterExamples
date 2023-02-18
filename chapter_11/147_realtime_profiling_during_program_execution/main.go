package main

import (
	"net/http"
	_ "net/http/pprof"
	"log"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8081", nil))
	}()

	// rest of the program
}