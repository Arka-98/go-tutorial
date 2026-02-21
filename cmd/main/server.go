package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"
)

// func init() {
// 	fmt.Println("Init example", os.Getenv("ENV_VAR_01"))
// }

func InitServer() {
	const port string = "0.0.0.0:8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})
	fmt.Println("Server listening on port:", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalln("Error starting HTTP server", err)
	}
}
