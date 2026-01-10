package main

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/Arka-98/go-tutorial/internal/dsa/linked_list"
	"github.com/Arka-98/go-tutorial/internal/generics"
	"github.com/Arka-98/go-tutorial/internal/interfaces"
	"github.com/Arka-98/go-tutorial/internal/structs"
	"github.com/Arka-98/go-tutorial/internal/utils"
)

//go:embed embed_example.txt
var outputContent string

func init() {
	loadEnvFile()
}

func main() {
	/*
		variables demo
	*/
	// const name = "Arkadipta"
	// const employeeId int = 101
	// job, age := "Developer", 27

	// fmt.Println(utils.GetQuote(), name, employeeId, job, age)

	// const rem = 22 / 7

	// fmt.Println(rem, reflect.TypeOf(rem))

	// /*
	// 	array & slices demo
	// */
	// numsArr := [5]int{0, 1, 2}
	// fmt.Println(numsArr, numsArr[2], reflect.TypeOf(numsArr))

	// for _, v := range numsArr {
	// 	fmt.Printf("Index: %d Value: %d\n", v)
	// }

	// arr1 := [3]int{0, 1, 2}
	// arr2 := [3]int{5, 1, 2}

	// fmt.Println(arr1 == arr2)

	// arr1Slice := numsArr[2:3]
	// // sliceWithMake := make([]int, 3)

	// fmt.Println(cap(arr1Slice), len(arr1Slice))

	// var arr3 []int

	// fmt.Println(arr3, reflect.TypeOf(arr3))

	// /*
	// 	maps
	// */
	// // map1 := map[string]int{"a": 5, "b": 10, "c": 15}

	// // var nilMap map[string]int

	// // nilMap["key1"] = 5

	// // fmt.Println("Nil map", nilMap)

	// map2 := make(map[string]map[string]int)

	// // map2["key1"] = 5

	// map2["key1"] = map[string]int{"a": 5, "b": 10, "c": 15}
	// map2["key2"] = map[string]int{"a": 3, "b": 6}

	// delete(map2, "key2")

	// fmt.Println("non existing key", map2["dawd"], map2["key2"])

	// next := createCounter()

	// fmt.Println(next())
	// fmt.Println(next())
	// fmt.Println(next())
	// fmt.Println(next())
	// fmt.Println(next())

	// fmt.Println(applyOperation(2, 5, add))

	// val, err := process(-10)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(val)

	// fmt.Println("logging panic res", panicDeferExample(10))

	// val, fn := getValAndFunc()

	// fmt.Println(val, fn())

	// variadicExample(arr1[:]...)

	// fmt.Println("outside variadic example", arr1)

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
	// consoleApp()
	// diceGame()

	// namesMapById := map[int]string{0: "Sam", 1: "Tom", 2: "Jerry"}

	// fmt.Println(namesMapById)

	// urlParsing()
	// buildUrl()
	// readerExample("ABCDEFG")
	// writerExample("this is a writer test\n", os.Stdout)
	// fmt.Println(writingToFile("this is a sample text"))
	// fmt.Println(bufferedFileWrite("this is buffered file write\n"))
	// fmt.Println(readingFromFile())
	// fmt.Println(openAndReadFromFile())
	// fmt.Println(outputContent)
	// fmt.Println(os.Getenv("ENV_VAR_01"))
	// loggerExample()
	// jsonExample()
	// slicesExample()

	// exStruct := IStruct{
	// 	value: "val",
	// 	// log: func() {
	// 	// 	fmt.Println("struct log")
	// 	// },
	// }

	// exStruct.Log()

	// typesExample(exStruct)

	// fmt.Println("Before goroutine")

	// go goroutinesExample()

	// fmt.Println("After goroutine")

	// for range 100000 {}

	// fmt.Println("After processing")

	// fmt.Println("begin heavy processing")

	// go heavyFn()
	// go heavyFn()

	// time.Sleep(30 * time.Second)

	// channelsExample()

	// bufferedChannels()

	// bufferedChannels2()

	// channelSync()

	// channelSyncWithClose()

	// sendOnlyChannel()

	// multiplexWithSelect()

	// multiplexWithFor()

	// multiplexWithTwoChannels()

	// contextExample()

	// testFn1()

	// timerExample()

	// tickerExample()

	// testFn2()

	// workerPoolExample()

	// problems.Process()

	// waitGroupExample()

	// problems.ProcessWg()

	// problems.MutexEx()

	// problems.AtomicExample()

	// testFn3()

	// fmt.Println(int(0.3/0.1))

	// problems.TokenBucketEx()

	// problems.FixedWindowEx()

	typeAliasAndNewTypeEx()
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

func panicDeferExample(input int) (myVar int) {
	myVar = 5

	defer fmt.Println("Deferred log 1", input)
	defer fmt.Printf("Deferred log 2 input = %d, myVar = %d\n", input, myVar)

	if input < 0 {
		panic("Negative numbers not allowed")
	}

	input = 30

	fmt.Printf("Normal log 1 input = %d, myVar = %d\n", input, myVar)
	defer fmt.Println("Deferred log 3", input)

	defer func() {
		myVar = 10

		fmt.Println("myVar is", myVar)
	}()

	return func() int {
		fmt.Println("returning")

		return myVar
	}()
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
			continue
		}

		err = parsedTemplates[optionMapKey].Execute(os.Stdout, dataMap)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func diceGame() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Menu")
		fmt.Println("1. Roll dice")
		fmt.Println("2. Exit")
		fmt.Println("Choose an option")

		option, _ := reader.ReadString('\n')

		switch strings.TrimSpace(option) {
		case "1":
			firstDie, secondDie := rollDie(1, 6), rollDie(1, 6)

			fmt.Printf("Die 1 - %d, Die 2 - %d\nTotal - %d\n", firstDie, secondDie, firstDie+secondDie)
		case "2":
			os.Exit(0)
		default:
			fmt.Println("Please choose the correct option")
		}
	}
}

func rollDie(lower, upper int) int {
	return rand.Intn(upper-lower+1) + lower
}

func urlParsing() {
	endpoint := "http://localhost:8001/api/v1/users/1?skip=0&limit=10&limit=15&limit=20"
	parsedUrl, err := url.Parse(endpoint)

	if err != nil {
		fmt.Println("Error parsing URL:", err)

		return
	}

	fmt.Println(parsedUrl.Host, parsedUrl.RawQuery, parsedUrl.Path, parsedUrl.Port())
	fmt.Println(parsedUrl.Query())
}

func buildUrl() {
	url := &url.URL{
		Scheme: "https",
		Host:   "localhost:4000",
		Path:   "api/v1/users",
	}
	query := url.Query()

	query.Set("skip", "10")
	query.Set("skip", "15")
	query.Set("limit", "20")

	url.RawQuery = query.Encode()

	fmt.Println("Built URL", url.String())

	str := "Hello golang"

	str = "da"

	fmt.Println(str)
}

// reads from provided str arg
func readerExample(str string) {
	reader := bufio.NewReader(strings.NewReader(str))
	data := make([]byte, 2)

	for {
		n, err := reader.Read(data)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading string:", err)

			return
		}

		fmt.Println(string(data[:n]))
	}
}

// writes provided str arg to a provided writer
func writerExample(str string, w io.Writer) {
	writer := bufio.NewWriter(w)
	n, err := writer.WriteString(str)

	if err != nil {
		fmt.Println("Error writing:", err)

		return
	}

	err = writer.Flush()

	if err != nil {
		fmt.Println("Error flushing:", err)

		return
	}

	fmt.Printf("Written %d bytes", n)
}

func writingToFile(text string) (err error) {
	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Error creating file:", err)

		return
	}

	_, err = file.Write([]byte(text))

	if err != nil {
		fmt.Println("Error writing to file:", err)

		return
	}

	defer func() {
		err = file.Close()
	}()

	return
}

func bufferedFileWrite(text string) (err error) {
	file, err := os.Create("output-2.txt")

	if err != nil {
		return
	}

	writer := bufio.NewWriter(file)

	_, err = writer.Write([]byte(text))

	if err != nil {
		return
	}

	defer func() {
		err = writer.Flush()
		err = file.Close()
	}()

	return
}

func readingFromFile() (err error) {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading from file:", err)

		return
	}

	fmt.Println(string(content))

	return
}

func openAndReadFromFile() (err error) {
	file, err := os.Open("input.txt")

	defer func() {
		err = file.Close()
	}()

	if err != nil {
		fmt.Println("Error reading from file:", err)

		return
	}

	reader := bufio.NewReader(file)
	data := make([]byte, 16)

	for {
		_, err = reader.Read(data)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading from file:", err)

			return
		}

		fmt.Println("Content:", string(data))
	}

	return
}

func loadEnvFile() (err error) {
	file, err := os.Open(".env")

	if err != nil {
		return
	}

	defer func() {
		err = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		kvStr := strings.Split(line, "=")

		os.Setenv(kvStr[0], kvStr[1])
	}

	err = scanner.Err()

	return
}

func loggerExample() {
	logger, err := utils.NewLogger()

	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	defer logger.Close()

	logger.Info("Initializing dependencies")
	logger.Info("Dependencies initialized")
	logger.Debug("Creating default users")
	logger.Info("Service started successfully")
}

func jsonExample() {
	type Person struct {
		FirstName string          `json:"first_name,omitempty"`
		LastName  string          `json:"last_name,omitempty"`
		Address   structs.Address `json:"address"`
	}

	person := Person{
		FirstName: "Arkadipta",
		LastName:  "Das",
		Address: structs.Address{
			Street:  "Baker Street",
			City:    "London",
			Country: "UK",
		},
	}

	data, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error while marshalling:", err)

		return
	}

	fmt.Println(string(data))

	var newPerson Person

	err = json.Unmarshal(data, &newPerson)

	if err != nil {
		fmt.Println("Error while unmarshalling", err)

		return
	}

	fmt.Println(newPerson, newPerson == person)
}

func pipeExample() (err error) {
	pr, pw := io.Pipe()

	go func() {
		pw.Write([]byte("Heyaa"))

		err = pw.Close()
	}()

	buf := new(bytes.Buffer)

	buf.ReadFrom(pr)
	fmt.Println(buf.String())

	return
}

func slicesExample() {
	bytes := []byte{'a', 'b'}

	bytes = append(bytes, 78)

	fmt.Println(string(bytes))

	dataMapWithNew := new(map[string]int)
	dataMapWithMake := make(map[string]int)

	// (*dataMapWithNew)["a"] = 5
	dataMapWithMake["a"] = 5

	fmt.Println(dataMapWithNew, dataMapWithMake)
}

type IExample interface {
	Log()
}

type IStruct struct {
	value string
	// log   func()
}

func (s IStruct) Log() {
	fmt.Println("Explicit struct log function")
}

func typesExample(data IExample) {
	fmt.Printf("Type is %T and the data value is %v\n", data, data)

	data = IExample(data)

	fmt.Printf("Again Type is %T and the data value is %v", data, data)
}

func goroutinesExample() {
	fmt.Println("Hello from goroutine")
}

func heavyFn() {
	for range 100000000000 {
	}

	fmt.Println("Complete!")
}

func linkedListExample() {
	llist := linked_list.NewLinkedList()

	llist.Append(10)
	llist.Append(5)
	llist.Append(8)

	err := llist.Traverse()

	llist.Shift()

	err = llist.Traverse()

	llist.Unshift(3)

	err = llist.Traverse()

	if err != nil {
		fmt.Println(err)

		return
	}
}

func channelsExample() {
	greeting := make(chan string)
	greeter := "Hello"

	go func() {
		fmt.Println("Init goroutine")

		greeting <- greeter

		fmt.Println("Sent to channel")
	}()

	fmt.Println("Outside goroutine")

	time.Sleep(3 * time.Second)

	fmt.Println("Before reciving")

	receiver := <-greeting

	fmt.Println("After receiving")
	fmt.Println(receiver)
}

// blocking receive from buffered channel
func bufferedChannels() {
	bufChannel := make(chan int, 2)

	bufChannel <- 5
	bufChannel <- 10

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("1st goroutine completed")
	}()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("2nd goroutine completed")
	}()

	fmt.Println("Received:", <-bufChannel)
	fmt.Println("Received:", <-bufChannel)
	fmt.Println("Received:", <-bufChannel)
	fmt.Println("End")
}

// blocking send from buffered channel
func bufferedChannels2() {
	bufChannel := make(chan int, 2)

	bufChannel <- 5
	bufChannel <- 10

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("1st goroutine completed", <-bufChannel)
	}()

	bufChannel <- 15

	fmt.Println("End")
}

func channelSync() {
	numGoroutines := 3
	data := make(chan int, numGoroutines)

	for i := range numGoroutines {
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println("from goroutine:", i)
			data <- i * 2
		}()
	}

	fmt.Println("Blocking receivers till value sent to channel")

	for range numGoroutines {
		fmt.Println("received:", <-data)
	}
}

func channelSyncWithClose() {
	ch := make(chan int)

	go func() {
		for i := range 3 {
			ch <- i * 2

			fmt.Println("Sent:", i)
		}

		close(ch)
	}()

	fmt.Println("Blocking till value is sent")

	for re := range ch {
		fmt.Println("Received:", re)
	}
}

func sendOnlyChannel() {
	ch := make(chan int)

	produce(ch)
	consume(ch)
}

func produce(ch chan<- int) {
	go func() {
		for i := range 5 {
			time.Sleep(800 * time.Millisecond)

			ch <- (i + 1) * 2
		}

		close(ch)
	}()
}

func consume(ch <-chan int) {
	for rcv := range ch {
		fmt.Println("Received:", rcv)
	}
}

func multiplexWithSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// close(ch1)
	// close(ch2)

	produce(ch1)
	produce(ch2)

	for range 2 {
		select {
		case val := <-ch1:
			fmt.Println("Ch 1 received:", val)
		case val := <-ch2:
			fmt.Println("Ch 2 received:", val)
			// default:
			// 	fmt.Println("Channels are not ready")
		}
	}

	fmt.Println("End")
}

func multiplexWithFor() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			time.Sleep(time.Second)

			ch <- (i + 1) * 2
		}

		close(ch)
	}()

loop:

	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")

				break loop
			}

			fmt.Println("Received:", data)
		}
	}

	// for re := range ch {
	// 	fmt.Println("Received:", re)
	// }

	fmt.Println("End")
}

func multiplexWithTwoChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := range 10 {
			time.Sleep(500 * time.Millisecond)

			ch1 <- i * 2
		}

		close(ch1)
	}()
	go func() {
		for re := range ch1 {
			fmt.Println("Ch1 received:", re)

			if re%2 == 0 {
				ch2 <- re
			}
		}

		close(ch2)
	}()

	for {
		select {
		case data, ok := <-ch2:
			if !ok {
				fmt.Println("Channel closed")

				return
			}

			fmt.Println("Ch2 received:", data)
		}
	}
}

func testFn1() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)

		close(ch)
	}()

	fmt.Println("Blocking starts...")

	fmt.Println(<-ch)

	// for re := range ch {
	// 	fmt.Println("Received:", re)
	// }
}

func testFn2() {
	bufCh := make(chan int, 3)

	go func() {
		for re := range bufCh {
			fmt.Println("Received on goroutine 1:", re)
		}

		fmt.Println("Complete goroutine 1")
	}()
	go func() {
		for re := range bufCh {
			fmt.Println("Received on goroutine 2:", re)
		}

		fmt.Println("Complete goroutine 2")
	}()

	fmt.Println("Sleep start")

	time.Sleep(2 * time.Second)

	fmt.Println("Sleep end")

	for i := range 5 {
		val := (i + 1) * 2
		bufCh <- val

		fmt.Println("Sent", val)
	}

	close(bufCh)

	time.Sleep(time.Second)

	fmt.Println("End")
}

func contextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	ctx = context.WithValue(ctx, "user", "jwt1234")

	time.Sleep(3 * time.Second)

	fmt.Println("Blocking starts")

	select {
	case _, ok := <-ctx.Done():
		fmt.Printf("Timed out. Channel closed: %v, Err: %v \n", !ok, ctx.Err())
	default:
		fmt.Println(ctx.Value("user"))
	}

	// <-ctx.Done()

	// for re := range ctx.Done() {
	// 	fmt.Println("Timed out, expired:", re, ctx.Err())
	// }

	fmt.Println("After:", ctx.Value("user"))
}

func timerExample() {
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Sleep starts")
	time.Sleep(1 * time.Second)
	fmt.Println("Sleep ends")

	stopped := timer.Stop()

	fmt.Println("Stopped:", stopped)

	// timer channel receive
	fmt.Println("timer receive starts")
	res := <-timer.C

	fmt.Printf("new timer ended: %v, Stopped: %v \n", res, stopped)
}

func tickerExample() {
	ticker := time.NewTicker(time.Second)

	for curTime := range ticker.C {
		fmt.Println(curTime)
	}
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d: Received job - %d\n", id, task)

		// Processing
		time.Sleep(time.Second)

		results <- task * 2
	}
}

func workerPoolExample() {
	numTasks := 10
	numWorkers := 3
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// initiate workers
	for i := range numWorkers {
		go worker(i+1, tasks, results)
	}

	// send tasks
	for i := range numTasks {
		tasks <- i + 1
	}

	close(tasks)

	// receive results
	for range numTasks {
		fmt.Println("Result - ", <-results)
	}
}

func waitGroupWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Initiated goroutine:", id)

	// Processing
	time.Sleep(time.Second)

	fmt.Println("Finished goroutine:", id)
}

func waitGroupExample() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)

		go waitGroupWorker(i+1, &wg)
	}

	wg.Wait()

	fmt.Println("End")
}

func testFn3() {
	bufCh := make(chan int)

	go func() {
		fmt.Println("goroutine sleep start..")

		time.Sleep(500 * time.Millisecond)

		for i := range 10 {
			bufCh <- i + 1

			time.Sleep(300 * time.Millisecond)
		}

		bufCh <- 11

		close(bufCh)
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("Waiting...")

	for re := range bufCh {
		fmt.Println("Received:", re)
	}
}

type EmployeeID int

func (e EmployeeID) isValid() bool {
	if e > 0 && e < 1000 {
		return true
	}

	return false
}

type TestFnType func(a, b int) bool

func (fn TestFnType) allow(a, b int) bool {
	return fn(a, b)
}

func typeAliasAndNewTypeEx() {
	empId := EmployeeID(1)
	testFn := func(e EmployeeID) {
		fmt.Printf("Type: %T, Value: %d, Reflect Type: %v, Valid: %v\n", e, e, reflect.TypeFor[EmployeeID](), e.isValid())
	}

	testFn(empId)
	fmt.Println(TestFnType(func(a, b int) bool {
		return a > b
	}).allow(11, 10))
}
