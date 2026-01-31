package concurrencycourse

import (
	"fmt"
	"sync"
	"time"
)

const HUNGER_COUNT = 3

type Fork struct {
	value int
	mu    sync.Mutex
}

func NewFork(value int) *Fork {
	return &Fork{value: value}
}

type Philosopher struct {
	name      string
	leftFork  *Fork
	rightFork *Fork
	hunger    int
}

func (ph *Philosopher) hasLeftFork() {
	fmt.Printf("%s has left fork %d. Time: %v\n", ph.name, ph.leftFork.value, time.Now())
}

func (ph *Philosopher) hasRightFork() {
	fmt.Printf("%s has right fork %d. Time: %v\n", ph.name, ph.rightFork.value, time.Now())
}

// When philosopher acquired both forks.
func (ph *Philosopher) eat() {
	ph.hunger--
	fmt.Printf("%s has both forks and is eating. Time: %v\n", ph.name, time.Now())
	// time.Sleep(2 * time.Second)
}

// When philosopher did not acquire any forks.
func (ph *Philosopher) think() {
	fmt.Printf("%s is thinking. Time: %v\n", ph.name, time.Now())
	// time.Sleep(500 * time.Millisecond)
}

func (ph *Philosopher) leave() {
	fmt.Printf("%s has left the table. Time: %v\n", ph.name, time.Now())
}

func createPhilosophers() []*Philosopher {
	philosopherNames := []string{
		"Newton",
		"Dijsktra",
		"Turing",
		"Einstein",
		"Maxwell",
	}
	philosophers := make([]*Philosopher, 0, len(philosopherNames))
	forks := make([]Fork, 0, len(philosopherNames))

	for i := range len(philosopherNames) {
		forks = append(forks, *NewFork(i))
	}

	for i, name := range philosopherNames {
		philosophers = append(philosophers, &Philosopher{
			name:      name,
			leftFork:  &forks[i],
			rightFork: &forks[(i+1)%len(philosopherNames)],
			hunger:    HUNGER_COUNT,
		})
	}

	return philosophers
}

func dine(philosopher *Philosopher, pwg *sync.WaitGroup, once *sync.Once, completionOrderCh chan<- string) {
	pwg.Done()
	pwg.Wait()
	once.Do(func() {
		fmt.Println("All philosophers are seated.")
	})

	for philosopher.hunger > 0 {
		if philosopher.leftFork.value < philosopher.rightFork.value {
			philosopher.leftFork.mu.Lock()
			philosopher.hasLeftFork()
			philosopher.rightFork.mu.Lock()
			philosopher.hasRightFork()

			philosopher.eat()

			philosopher.leftFork.mu.Unlock()
			philosopher.rightFork.mu.Unlock()
		} else {
			philosopher.rightFork.mu.Lock()
			philosopher.hasRightFork()
			philosopher.leftFork.mu.Lock()
			philosopher.hasLeftFork()

			philosopher.eat()

			philosopher.rightFork.mu.Unlock()
			philosopher.leftFork.mu.Unlock()
		}

		philosopher.think()
	}

	philosopher.leave()
	completionOrderCh <- philosopher.name
}

func DiningPhilosopherEx() {
	var (
		pwg  sync.WaitGroup
		ewg sync.WaitGroup
		wg   sync.WaitGroup
		once sync.Once
	)

	completionOrderCh := make(chan string)
	philosophers := createPhilosophers()
	completionOrder := make([]string, 0, len(philosophers))

	pwg.Add(len(philosophers))
	ewg.Go(func() {
		for philosopherName := range completionOrderCh {
			completionOrder = append(completionOrder, philosopherName)
		}
	})

	for i := range philosophers {
		wg.Go(func() {
			dine(philosophers[i], &pwg, &once, completionOrderCh)
		})
	}

	wg.Wait()
	close(completionOrderCh)
	ewg.Wait()

	fmt.Println(completionOrder)
}
