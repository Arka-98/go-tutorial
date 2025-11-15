package main

import (
	"fmt"
	"reflect"
	// "github.com/Arka-98/go-tutorial/internal/problems"
	"github.com/Arka-98/go-tutorial/internal/utils"
)

func main() {
	const name = "Arkadipta"
	const employeeId = 101
	job, age := "Developer", 27

	fmt.Println(utils.GetQuote(), name, employeeId, job, age)

	const rem = 22 / 7

	fmt.Println(rem, reflect.TypeOf(rem))

	// problems.StarPattern(5)
}
