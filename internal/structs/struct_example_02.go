package structs

import "fmt"

type Bird struct {
	Name string
}

type Dinosaur struct {
	Class string
}

type Whale struct {
	Name string
}

// Eat: Bird methods.
func (bird Bird) Eat(food string) string {
	return fmt.Sprintf("%v eats %v!", bird.Name, food)
}

func (bird Bird) Sleep() string {
	return fmt.Sprintf("%v sleeps!", bird.Name)
}

// Dinosaur methods
func (dinosaur Dinosaur) Eat(food string) string {
	return fmt.Sprintf("%v eats %v!", dinosaur.Class, food)
}

func (dinosaur Dinosaur) Sleep() string {
	return fmt.Sprintf("%v sleeps!", dinosaur.Class)
}

func (dinosaur Dinosaur) Destroy() string {
	return fmt.Sprintf("%v destroys!", dinosaur.Class)
}

// Whale methods
func (whale Whale) Eat(food string) string {
	return fmt.Sprintf("%v eats %v!", whale.Name, food)
}

func (whale Whale) Yawn() string {
	return fmt.Sprintf("%v Yawn!", whale.Name)
}
