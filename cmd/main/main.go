package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"os"
	"reflect"
	"strings"

	"github.com/Arka-98/go-tutorial/internal/generics"
	"github.com/Arka-98/go-tutorial/internal/interfaces"

	// "github.com/Arka-98/go-tutorial/internal/problems"

	// "github.com/Arka-98/go-tutorial/internal/problems"
	"github.com/Arka-98/go-tutorial/internal/structs"
	"github.com/Arka-98/go-tutorial/internal/utils"
)

func main() {
	/*
		variables demo
	*/
	const name = "Arkadipta"
	const employeeId int = 101
	job, age := "Developer", 27

	fmt.Println(utils.GetQuote(), name, employeeId, job, age)

	const rem = 22 / 7

	fmt.Println(rem, reflect.TypeOf(rem))

	/*
		array & slices demo
	*/
	numsArr := [5]int{0, 1, 2}
	fmt.Println(numsArr, numsArr[2], reflect.TypeOf(numsArr))

	for _, v := range numsArr {
		fmt.Printf("Index: %d Value: %d\n", v)
	}

	arr1 := [3]int{0, 1, 2}
	arr2 := [3]int{5, 1, 2}

	fmt.Println(arr1 == arr2)

	arr1Slice := numsArr[2:3]
	// sliceWithMake := make([]int, 3)

	fmt.Println(cap(arr1Slice), len(arr1Slice))

	var arr3 []int

	fmt.Println(arr3, reflect.TypeOf(arr3))

	/*
		maps
	*/
	// map1 := map[string]int{"a": 5, "b": 10, "c": 15}

	// var nilMap map[string]int

	// nilMap["key1"] = 5

	// fmt.Println("Nil map", nilMap)

	map2 := make(map[string]map[string]int)

	// map2["key1"] = 5

	map2["key1"] = map[string]int{"a": 5, "b": 10, "c": 15}
	map2["key2"] = map[string]int{"a": 3, "b": 6}

	delete(map2, "key2")

	fmt.Println("non existing key", map2["dawd"], map2["key2"])

	next := createCounter()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	fmt.Println(applyOperation(2, 5, add))

	val, err := process(-10)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(val)

	panicDeferExample(10)

	val, fn := getValAndFunc()

	fmt.Println(val, fn())

	variadicExample(arr1[:]...)

	fmt.Println("outside variadic example", arr1)

	// problems.StarPattern(5)

	// array of map values
	// createArrayOfMapElems()

	// var arr4 []int
	// i, j := 5, 10
	// sum := 15
	// slice1 := []int{i, j}

	// fmt.Println(arr4)

	// arr4 = append(arr4, 5)

	// fmt.Println(arr4)

	// if i != 0 {

	// }

	// fmt.Println(problems.Factorial(5))
	// fmt.Println(problems.Fibonacci(5))
	// fmt.Println(problems.SumOfDigits(986))

	// pointersExample()
	// runeExample()
	// structExample()
	// interfaceExample()
	// fmt.Println(genericsExample([]int{1, 2, 3}))
	// genericsExample()
	// errorHandling()
	// textTemplates()

	consoleApp()
}

func add(operand1, operand2 int) int {
	return operand1 + operand2
}

func createCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func applyOperation(operand1, operand2 int, operation func(int, int) int) int {
	return operation(operand1, operand2)
}

func process(val int) (int, error) {
	if val < 0 {
		return 0, errors.New("Negative values are not supported")
	}

	return val, nil
}

func panicDeferExample(input int) {
	defer fmt.Println("Deferred log 1", input)
	// input = 20
	defer fmt.Println("Deferred log 2", input)

	if input < 0 {
		panic("Negative numbers not allowed")
	}

	input = 30

	fmt.Println("Normal log 1", input)
	defer fmt.Println("Deferred log 3", input)
}

func getValAndFunc() (int, func() int) {
	return 10, func() int {
		return 50
	}
}

func variadicExample(vals ...int) {
	vals[2] = 10

	fmt.Println("variadic example", vals)
}

func createArrayOfMapElems() {
	arrOfMaps := make(map[string][]int)

	arrOfMaps["key1"] = make([]int, 5)
	arrOfMaps["key1"] = append(arrOfMaps["key1"], 10)

	arrOfMaps["key2"] = []int{0, 1, 2}

	fmt.Println(arrOfMaps)

	mapArr := []map[string]int{{"m11": 5, "m12": 10}, {"m21": 6, "m22": 12}}

	fmt.Println(mapArr)
}

func pointersExample() {
	var ptr *int

	x := 10

	ptr = &x

	fmt.Println(x, ptr, *ptr)

	*ptr++

	fmt.Println(x, *ptr)

	var doublePtr **int
	var singlePtr *int

	y := 10
	// val2 := 10

	singlePtr = &y
	doublePtr = &singlePtr

	*singlePtr++
	**doublePtr++

	fmt.Println(*singlePtr, **doublePtr)
}

func runeExample() {
	str := "I am a developer"
	rn := 'A'

	fmt.Println(string(rn), str)

	for _, v := range str {
		fmt.Printf("Type: %T Value: %v\n", v, string(v))
	}
}

func structExample() {
	addr := structs.Address{
		Street:  "Kwality",
		City:    "Kolkata",
		Country: "India",
	}
	org := structs.Organization{
		ID:      1,
		Name:    "HCL Software",
		Ceo:     "John Doe",
		Address: addr,
	}
	emp := structs.Employee{
		ID:   1,
		Name: "Arkadipta",
		Role: "Senior Software Engineer",
		Org:  org,
		Address: structs.Address{
			Street:  "Park Street",
			City:    "Manhattan",
			Country: "US",
		},
	}

	addrNew := structs.Address{"dwa", "d", "a"}

	fmt.Println(emp, addrNew)
	fmt.Println(emp.GetEmployeeDetails())
}

func interfaceExample() {
	vulture := structs.Bird{
		Name: "BrainEater",
	}
	trex := structs.Dinosaur{
		Class: "T-Rex",
	}
	// freddie := structs.Whale{
	// 	Name: "Freddie The Whale",
	// }

	animalExample(vulture, "Meat")
	animalExample(trex, "Meat")
	// animalExample(freddie, "Octopus")
}

func animalExample(animal interfaces.Animal, food string) {
	fmt.Println(animal.Eat(food))
	fmt.Println(animal.Sleep())
}

func genericsExample() {
	fmt.Println(generics.Add(5, 3))

	empErr := generics.ParseResponse1(generics.Response[generics.Employee]{
		Data: generics.Employee{
			ID:   101,
			Name: "Jack",
			Job:  "SSE",
		},
		Errored: false,
	})

	if empErr != nil {
		fmt.Println(empErr)
	}

	projErr := generics.ParseResponse1(generics.Response[generics.Project]{
		Data: generics.Project{
			ID:     901,
			Title:  "HCL Commerce",
			Client: "HCLSoftware",
		},
		Errored: false,
	})

	if projErr != nil {
		fmt.Println(projErr)
	}
}

func errorHandling() {
	res, err := structs.FindId(10)

	if err != nil {
		if _, ok := err.(*structs.CustomError); ok {
			fmt.Println("Custom Error caught", err)
		} else {
			fmt.Println("Error", err)
		}

		return
	}

	fmt.Println(res)
}

func textTemplates() {
	tmpl, err := template.New("example").Parse("Hey {{.name}}, working as {{.job}}")

	if err != nil {
		fmt.Println(err)

		return
	}

	err = tmpl.Execute(os.Stdout, map[string]string{"name": "Arka", "job": "SDE"})

	if err != nil {
		fmt.Println(err)

		return
	}
}

func consoleApp() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	templatesMapByName := map[string]string{
		"welcome":      "Hi {{.name}}. Welcome here!",
		"notification": "{{.name}} you have a new notification. {{.notification}}",
		"error":        "Something went wrong. {{.error}}",
	}

	parsedTemplates := make(map[string]*template.Template)

	for name, tmpl := range templatesMapByName {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Join")
		fmt.Println("2. Get notification")
		fmt.Println("3. Get error")
		fmt.Println("4. Quit")
		fmt.Println("Choose your option")

		option, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)

			return
		}

		var optionMapKey string
		dataMap := map[string]string{"name": name}

		switch strings.TrimSpace(option) {
		case "1":
			optionMapKey = "welcome"
		case "2":
			optionMapKey = "notification"

			notification, _ := reader.ReadString('\n')

			dataMap[optionMapKey] = notification
		case "3":
			optionMapKey = "error"

			err, _ := reader.ReadString('\n')

			dataMap[optionMapKey] = err
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Please choose the correct option")
		}

		if optionMapKey == "" {
			continue
		}

		err = parsedTemplates[optionMapKey].Execute(os.Stdout, dataMap)

		if err != nil {
			fmt.Println(err)
		}
	}
}
