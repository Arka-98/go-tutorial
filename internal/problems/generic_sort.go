package problems

import (
	"fmt"
	"sort"
)

type Employee struct {
	name   string
	age    int
	salary int
	role   string
}

type By func(e1, e2 *Employee) bool

type EmployeeSorter struct {
	emps []Employee
	by   By
}

func (e *EmployeeSorter) Len() int {
	return len(e.emps)
}

func (e *EmployeeSorter) Less(i, j int) bool {
	return e.by(&e.emps[i], &e.emps[j])
}

func (e *EmployeeSorter) Swap(i, j int) {
	e.emps[i], e.emps[j] = e.emps[j], e.emps[i]
}

func (by By) Sort(emps []Employee) {
	empSorter := &EmployeeSorter{
		emps: emps,
		by:   by,
	}

	sort.Sort(empSorter)
}

// type EmployeesByAge []Employee

// func (e EmployeesByAge) Len() int {
// 	return len(e)
// }

// func (e EmployeesByAge) Less(i, j int) bool {
// 	return e[i].age < e[j].age
// }

// func (e EmployeesByAge) Swap(i, j int) {
// 	e[i], e[j] = e[j], e[i]
// }

func GenericSortEx() {
	// intSlice := []int{10, 3, 8, 7, 2}

	// sort.Ints(intSlice)
	// fmt.Println(intSlice)

	// emps := EmployeesByAge([]Employee{
	// 	{"Sam", 32, 50000, "SSE1"},
	// 	{"John", 27, 180000, "SSE2"},
	// 	{"Parek", 25, 200000, "SSE3"},
	// })

	// sort.Sort(emps)
	// fmt.Println(emps)

	emps := []Employee{
		{"Sam", 32, 50000, "SSE1"},
		{"John", 27, 180000, "SSE2"},
		{"Parek", 25, 200000, "SSE3"},
	}

	// var empsByAge By = func(e1, e2 *Employee) bool {
	// 	return e1.age < e2.age
	// }
	var empsBySalary By = func(e1, e2 *Employee) bool {
		return e1.salary < e2.salary
	}

	fmt.Println("before", emps)

	By(empsBySalary).Sort(emps)

	fmt.Println("after", emps)
}
