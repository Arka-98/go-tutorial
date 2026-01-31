package problems

import (
	"fmt"
	"reflect"
	"sync"
)

type employeeRole string

const (
	se  employeeRole = "Software Engineer"
	sse employeeRole = "Senior Software Engineer"
	lse employeeRole = "Lead Software Engineer"
)

type employee struct {
	name string
	role employeeRole
}

func SyncPoolEx() {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("Creating new instance")

			return &employee{}
		},
	}

	emp1 := pool.Get().(*employee)
	emp1.name = "John"
	emp1.role = se

	fmt.Println(emp1, reflect.TypeOf(emp1))

	pool.Put(emp1)

	emp2 := pool.Get().(*employee)
	emp3 := pool.Get().(*employee)

	fmt.Println(emp2, emp3, reflect.TypeOf(emp2), reflect.TypeOf(emp3))
}
