package tests

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"

	concurrencycourse "github.com/Arka-98/go-tutorial/internal/concurrency_course"
)

func TestConcurrentAdd(t *testing.T) {
	var wg sync.WaitGroup

	r, w, err := os.Pipe()
	stdout := os.Stdout

	if err != nil {
		t.Errorf("Error creating os.Pipe: %v", err)
	}

	os.Stdout = w

	wg.Go(func() {
		concurrencycourse.Add(1)

		err = w.Close()

		if err != nil {
			fmt.Println("Error closing os writer:", err)
		}
	})
	wg.Wait()
	
	res, err := io.ReadAll(r)
	os.Stdout = stdout

	if err != nil {
		t.Errorf("Error reading from os reader: %v", err)
	}

	output := string(res)

	if !strings.Contains(output, "1") {
		t.Errorf("Expected %v but got %v", 1, output)
	}
}
