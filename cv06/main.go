package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		_, err = fmt.Fprintf(w, "Hello, this is Vojtěch Voleman (%s)\n", hostname)
		if err != nil {
			println(err.Error())
			return
		}
		println("Request from " + r.RemoteAddr)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	fmt.Printf("Server is listening on port 8080\n")
}
