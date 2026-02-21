package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Init example", os.Getenv("ENV_VAR_01"))
}

// func main() {
// 	ch := make(chan int)

// 	ch <- 5
// }