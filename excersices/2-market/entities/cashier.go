package entities

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Cashier struct {
	User
	Mutex *sync.Mutex
}

func (cs *Cashier) ServeClient(c *Client, ts chan *Transaction, m *Market, wg *sync.WaitGroup) {
	defer wg.Done()
	maxWait := 800
	minWait := 200

	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*600)

	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Failed transaction (timeout)")
		return
	case <-time.After(time.Millisecond * (time.Duration(rand.UintN(uint(maxWait-minWait)) + uint(minWait)))):
	}

	cs.Mutex.Lock()

	fmt.Printf("Cashier (%s) serving (%s)\n", cs.Name, c.Name)

	ts <- &Transaction{
		Items: c.Inventory,
		Date:  time.Now(),
	}

	cs.Mutex.Unlock()

	c.Leave(m)
}
