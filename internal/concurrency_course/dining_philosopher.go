package concurrencycourse

import (
	"fmt"
	"sync"
	"time"
)

const hungerCount = 3
const eatDuration = 2 * time.Second
const thinkDuration = 500 * time.Millisecond

// type Fork struct {
// 	value int
// 	mu    sync.Mutex
// }

type request struct {
	philosopher *Philosopher
	ack         chan struct{}
}

type Philosopher struct {
	name       string
	hunger     int
	leftFork   chan struct{}
	rightFork  chan struct{}
	doneCh     chan struct{}
	requestsCh chan request
}

func (ph *Philosopher) eat() {
	fmt.Printf("%s has both forks and is eating. Time: %v\n", ph.name, time.Now())
	time.Sleep(eatDuration)
}

func (ph *Philosopher) think() {
	fmt.Printf("%s is thinking. Time %v\n", ph.name, time.Now())
	time.Sleep(thinkDuration)
}

func (ph *Philosopher) leave() {
	fmt.Printf("%s has left the table. Time: %v\n", ph.name, time.Now())
}

func createPhilosophers(philosopherNames []string, requestsCh chan request, doneCh chan struct{}) []*Philosopher {
	philosophers := make([]*Philosopher, 0, len(philosopherNames))
	forks := make([]chan struct{}, len(philosopherNames))

	for i := range len(philosopherNames) {
		forks[i] = make(chan struct{}, 1)
		forks[i] <- struct{}{}
	}

	for i, name := range philosopherNames {
		philosophers = append(philosophers, &Philosopher{
			name:       name,
			hunger:     hungerCount,
			requestsCh: requestsCh,
			leftFork:   forks[i],
			rightFork:  forks[(i+1)%len(philosopherNames)],
			doneCh:     doneCh,
		})
	}

	return philosophers
}

func Coordinator(philosophers []*Philosopher, requestsCh chan request, doneCh chan struct{}) {
	var (
		next            request
		completedEating int
	)

	totalHungerCount := hungerCount * len(philosophers)
	queue := make([]request, 0, totalHungerCount)

	for {
		if len(queue) > 0 {
			next = queue[0]
		}

		select {
		case req := <-requestsCh:
			queue = append(queue, req)
		case next.ack <- struct{}{}:
			queue = queue[1:]
		case <-doneCh:
			completedEating++

			if completedEating == totalHungerCount {
				return
			}
		}
	}
}

func Dine(philosopher *Philosopher, completionOrderCh chan<- string) {
	for i := range philosopher.hunger {
		ack := make(chan struct{})
		philosopher.requestsCh <- request{philosopher: philosopher, ack: ack}

		// wait for permission from coordinator
		<-ack

		// try to take forks to start eating
		<-philosopher.leftFork
		<-philosopher.rightFork

		philosopher.eat()

		// put back the forks
		philosopher.leftFork <- struct{}{}
		philosopher.rightFork <- struct{}{}
		philosopher.doneCh <- struct{}{}

		if i == philosopher.hunger-1 {
			close(philosopher.doneCh)

			completionOrderCh <- philosopher.name
		}

		philosopher.think()
	}

	philosopher.leave()
}

func Initialize(
	philosopherNames []string,
	completionOrder *[]string,
	pwg *sync.WaitGroup,
	ewg *sync.WaitGroup,
) ([]*Philosopher, chan<- string) {
	completionOrderCh := make(chan string)
	requestsCh := make(chan request)
	doneCh := make(chan struct{})
	philosophers := createPhilosophers(philosopherNames, requestsCh, doneCh)

	ewg.Go(func() {
		Coordinator(philosophers, requestsCh, doneCh)
	})
	pwg.Add(len(philosophers))
	ewg.Go(func() {
		for philosopherName := range completionOrderCh {
			*completionOrder = append(*completionOrder, philosopherName)
		}
	})

	return philosophers, completionOrderCh
}

func DiningPhilosopherEx() {
	var (
		wg  sync.WaitGroup
		pwg sync.WaitGroup
		ewg sync.WaitGroup
	)

	philosopherNames := []string{
		"Newton",
		"Dijsktra",
		"Turing",
		"Einstein",
		"Maxwell",
	}
	completionOrder := make([]string, 0, len(philosopherNames))
	philosophers, completionOrderCh := Initialize(philosopherNames, &completionOrder, &pwg, &ewg)

	for _, philosopher := range philosophers {
		wg.Go(func() {
			Dine(philosopher, completionOrderCh)
		})
	}

	wg.Wait()
	close(completionOrderCh)
	ewg.Wait()

	fmt.Println(completionOrder)
}
