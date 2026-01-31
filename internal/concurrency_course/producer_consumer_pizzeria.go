package concurrencycourse

import (
	"fmt"
	"sync"
	"time"
)

type PizzaOrder struct {
	customerId   int
	orderMessage string
}

type OrderRes struct {
	customerId int
	success    bool
}

type Pizzeria struct {
	data   chan PizzaOrder
	result chan OrderRes
}

func NewPizzeria(numOrders int) *Pizzeria {
	return &Pizzeria{
		data:   make(chan PizzaOrder),
		result: make(chan OrderRes),
	}
}

func (p *Pizzeria) PlaceOrder(orders []PizzaOrder) {
	for _, order := range orders {
		p.data <- order

		fmt.Printf("Customer: %d. Order: %s received. %v\n", order.customerId, order.orderMessage, time.Now())
	}

	close(p.data)
}

func (p *Pizzeria) Process() {
	for order := range p.data {
		time.Sleep(time.Second)

		p.result <- OrderRes{order.customerId, true}
	}
}

func GeneratePizzaOrdersAndOrderMapByCustId(orderMessages []string) ([]PizzaOrder, map[int]string) {
	pizzaOrders := make([]PizzaOrder, 0, len(orderMessages))
	orderMessageMapByCustomerId := make(map[int]string)

	for i, orderMessage := range orderMessages {
		pizzaOrders = append(pizzaOrders, PizzaOrder{i + 1, orderMessage})
		orderMessageMapByCustomerId[i+1] = orderMessage
	}

	return pizzaOrders, orderMessageMapByCustomerId
}

func PizzeriaEx() {
	const numOrders = 10

	const numWorkers = 5

	var wg sync.WaitGroup

	pizzeria := NewPizzeria(numOrders)

	for range numWorkers {
		wg.Go(func() {
			pizzeria.Process()
		})
	}

	go func() {
		wg.Wait()

		close(pizzeria.result)
	}()

	pizzaOrders, orderMessageMapByCustomerId := GeneratePizzaOrdersAndOrderMapByCustId([]string{
		"Olives",
		"Tomato",
		"Cheese",
		"Chicken",
		"Salami",
		"Pepperoni",
		"Pork",
		"Bell Peppers",
		"Onion",
		"Broccoli",
	})

	go pizzeria.PlaceOrder(pizzaOrders)

	for res := range pizzeria.result {
		fmt.Printf(
			"Customer: %d. Order: %s, served. %v\n", res.customerId, orderMessageMapByCustomerId[res.customerId], time.Now(),
		)
	}
}
