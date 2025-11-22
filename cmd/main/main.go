package main

import (
	"errors"
	"fmt"
	"reflect"

	// "github.com/Arka-98/go-tutorial/internal/problems"
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
	createArrayOfMapElems()

	// var arr4 []int
	// i, j := 5, 10
	// sum := 15
	// slice1 := []int{i, j}

	// fmt.Println(arr4)

	// arr4 = append(arr4, 5)

	// fmt.Println(arr4)

	// if i != 0 {

	// }
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
