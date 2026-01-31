package concurrencycourse

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func GoroutineWgEx() {
	var wg sync.WaitGroup

	greetings := []string{"hello universe", "hello cosmos", "hello world"}

	for _, greeting := range greetings {
		wg.Go(func() {
			updateMessage(greeting)
			printMessage()
		})
	}

	wg.Wait()

	fmt.Println("End")
}
