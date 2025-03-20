package entities

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Cashier struct {
	User
	Mutex *sync.Mutex
}

func (cs *Cashier) ServeClient(c *Client, ts chan *Transaction) {
	maxWait := 800
	minWait := 200
	cs.Mutex.Lock()

	fmt.Printf("Cashier (%s) serving (%s)\n", cs.Name, c.Name)

	time.Sleep(time.Millisecond * (time.Duration(rand.UintN(uint(maxWait-minWait)) + uint(minWait))))

	ts <- &Transaction{
		Items: c.Inventory,
		Date:  time.Now(),
	}

	cs.Mutex.Unlock()

	c.Leave()
}
