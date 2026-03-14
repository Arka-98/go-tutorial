package structs

import (
	"fmt"
	"reflect"
)

type ValueGetter interface {
	getValue() string
}

type A struct {
	name   string
	valueA string
}

func (a *A) getValue() string {
	return a.valueA
}

type B struct {
	// name   string
	valueB string
}

func (b *B) getValue() string {
	return b.valueB
}

type C struct {
	A
	B

	name string
	valueC string
}

type D struct {
	ValueGetter

	valueD string
}

func (c *C) getValue() string {
	return c.valueC
}

func parseValue(getter ValueGetter) {
	if reflect.ValueOf(getter.getValue()).IsZero() {
		fmt.Println("zero value")
	} else {
		fmt.Println("has value")
	}
}

func EmbedExample() {
	c := &C{
		name: "C name",
		valueC: "C value",
		A: A{
			name:   "A name",
			valueA: "A value",
		},
		B: B{
			// name:   "B name",
			valueB: "B value",
		},
	}
	d := &D{
		valueD: "D value",
		ValueGetter: c,
	}

	fmt.Println(d.getValue(), d.ValueGetter.getValue())

	parseValue(c)
	parseValue(d)
}
