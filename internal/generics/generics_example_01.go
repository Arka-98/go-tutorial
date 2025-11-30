package generics

import "fmt"

type Employee struct {
	ID   int
	Name string
	Job  string
}

type Project struct {
	ID     int
	Title  string
	Client string
}

type Response[T any] struct {
	Data    T
	Errored bool
}

func Add[T int | float32](num1, num2 T) T {
	return num1 + num2
}

func ParseResponse1[T Project | Employee](res Response[T]) error {
	if res.Errored {
		return fmt.Errorf("Error!")
	}

	var data any

	fmt.Printf("Type of data before assignment %T\n", data)

	switch v := any(res.Data).(type) {
		case Employee:
			data = v.Name
		case Project:
			data = v.Title
	}

	fmt.Printf("Type of data after assignment %T\n", data)
	fmt.Println(data)

	return nil
}