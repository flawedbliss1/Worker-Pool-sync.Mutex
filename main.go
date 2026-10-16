package main

import (
	"fmt"
	"sync"
)

type Evento struct {
	Nome string
	Acao string
}

type Counter struct {
	mu       sync.Mutex
	contador int
}

func (c *Counter) Incrementar() {
	c.mu.Lock()
	c.contador++
	c.mu.Unlock()
}

func worker(ch chan Evento, wg *sync.WaitGroup, c *Counter) {
	defer wg.Done()
	for eventos := range ch {
		c.Incrementar()
		fmt.Printf("Nome: %d %s %s \n", c.contador, eventos.Nome, eventos.Acao)
	}
}

func main() {
	evento := []Evento{
		{Nome: "alice", Acao: "entrou"},
		{Nome: "andrews", Acao: "entrou"},
		{Nome: "alice", Acao: "saiu"},
		{Nome: "andrews", Acao: "entrou"},
		{Nome: "eu", Acao: "entrei"},
	}
	ch := make(chan Evento)
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers)

	var c Counter

	for i := 0; i < numWorkers; i++ {
		go worker(ch, &wg, &c)
	}

	for _, j := range evento {
		ch <- j
	}
	close(ch)
	wg.Wait()
	fmt.Println("Total:", c.contador)
}
