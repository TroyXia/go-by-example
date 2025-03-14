package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Server() {
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		// flusher := w.(http.Flusher)
		for i := 0; i < 2; i++ {
			fmt.Fprintf(w, "huyun troy\n")
			// flusher.Flush()
			<-time.Tick(1 * time.Second)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	Server()
}
