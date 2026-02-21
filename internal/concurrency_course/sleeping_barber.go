// This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem
// which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
// a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is
// open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the
// waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to
// sleep until a new client arrives. So the rules are as follows:
//
//   - if there are no customers, the barber falls asleep in the chair
//   - a customer must wake the barber if he is asleep
//   - if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and
//     sits in an empty chair if it's available
//   - when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers
//     and falls asleep if there are none
//   - shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is
//     empty
//   - after the shop is closed and there are no clients left in the waiting area, the barber
//     goes home
//
// The Sleeping Barber was originally proposed in 1965 by computer science pioneer Edsger Dijkstra.
//
// The point of this problem, and its solution, was to make it clear that in a lot of cases, the use of
// semaphores (mutexes) is not needed.
package concurrencycourse

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

const (
	cutHairDuration       = 3 * time.Second
	shopOpenDuration      = 15 * time.Second
	clientArrivalDuration = 1 * time.Second
	shopBuffer            = 3
	numBarbers            = 2
)

type client struct {
	name string
}

func (c *client) cutHair(barberId int) {
	color.Yellow("Barber %d: %s is cutting his hair. Time: %v\n", barberId, c.name, time.Now())
	time.Sleep(cutHairDuration)
	color.Green("Barber %d: %s is done cutting his hair. Time: %v\n", barberId, c.name, time.Now())
}

func closeShop(clientsCh chan client, shopClosedCh chan<- struct{}) {
	time.Sleep(shopOpenDuration)
	color.Red("Closing barber shop for the day. Time: %v", time.Now())
	close(clientsCh)

	shopClosedCh <- struct{}{}
}

func barber(id int, clientsCh <-chan client) {
	for client := range clientsCh {
		client.cutHair(id)
	}
}

func produceClients(clientsCh chan<- client, shopClosedCh <-chan struct{}, clients []client) {
	for _, client := range clients {
		select {
		case <-shopClosedCh:
			color.Red("Barber shop is closed. Exiting.")

			return
		case clientsCh <- client:
			color.Cyan("%s takes a seat at the barbershop", client.name)

		default:
			color.Red("Barber is busy. Client %s is leaving. Time: %v\n", client.name, time.Now())
		}

		time.Sleep(clientArrivalDuration)
	}
}

func SleepingBarberEx() {
	var (
		produceClientsWg sync.WaitGroup
		barberWg         sync.WaitGroup
	)

	clientsCh := make(chan client, shopBuffer)
	shopClosedCh := make(chan struct{})
	clients := []client{
		{"John"},
		{"Eva"},
		{"Raimi"},
		{"Sam"},
		{"Tom"},
		{"Jamie"},
		{"Arthur"},
		{"Richard"},
		{"James"},
		{"Clark"},
		{"Bush"},
		{"Bob"},
		{"Dexter"},
		{"Morty"},
		{"Dan"},
	}

	go closeShop(clientsCh, shopClosedCh)

	produceClientsWg.Go(func() {
		produceClients(clientsCh, shopClosedCh, clients)
	})

	for i := range numBarbers {
		barberWg.Go(func() {
			barber(i, clientsCh)
		})
	}

	barberWg.Wait()
	produceClientsWg.Wait()
}
