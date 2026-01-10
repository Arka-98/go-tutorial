package problems

import (
	"fmt"
	"time"
)

type TicketRequest struct {
	personId   int
	numTickets int
	cost       int
}

type TicketResponse struct {
	PersonId int
	Status   bool
}

func ticketProcessor(id int, requests <-chan TicketRequest, results chan<- TicketResponse) {
	for request := range requests {
		fmt.Printf("Worker: %d, Received person ID: %d\n", id, request.personId)

		// Processing
		time.Sleep(time.Duration(request.numTickets) * time.Second)

		results <- TicketResponse{PersonId: request.personId, Status: true}
	}
}

func Process() {
	numPersons := 10
	numWorkers := 3
	requests := make(chan TicketRequest, numPersons)
	results := make(chan TicketResponse, numPersons)

	for i := range numWorkers {
		go ticketProcessor(i+1, requests, results)
	}

	for i := range numPersons {
		var numTickets int

		if i%2 == 0 {
			numTickets = 2
		} else {
			numTickets = 1
		}

		requests <- TicketRequest{personId: i + 1, numTickets: numTickets, cost: 100 * numTickets}
	}

	close(requests)

	for range numPersons {
		ticketRes := <-results

		fmt.Printf("Ticket booking status for person ID %d: %v\n", ticketRes.PersonId, ticketRes.Status)
	}
}
