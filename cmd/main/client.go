package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func clientExample() {
	client := &http.Client{}

	res, err := client.Get("https://jsonplaceholder.typicode.com/todos/1/")

	if err != nil {
		log.Println("Error making GET request:", err)

		return
	}

	defer func() {
		err = res.Body.Close()

		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)

		return
	}

	fmt.Println(string(body))
}