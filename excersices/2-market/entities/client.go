package entities

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sync"
	"time"
)

type ClientInventory struct {
	Product
	Quantity uint
}

type Client struct {
	User
	ID        uint
	Inventory []*ClientInventory
}

func (c *Client) PickupProducts(m *Market, q chan *Client, wg *sync.WaitGroup) {
	defer wg.Done()
	const maxPickup = 4
	const maxTime = 200
	pickupCount := rand.UintN(maxPickup-1) + 1

	for i := uint(0); i < pickupCount; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(maxTime)))

		if len(m.Inventories) == 0 {
			return
		}

		idx := rand.IntN(len(m.Inventories))
		prod := m.Inventories[idx]

		prod.Mutex.Lock()
		if prod.Stock < 1 {
			prod.Mutex.Unlock()
			continue
		}

		qty := rand.UintN(maxPickup-1) + 1
		if qty > prod.Stock {
			qty = prod.Stock
		}

		fmt.Printf("Client (%s) picking up Product: (%s), Qty: (%d) Price: (%f)\n", c.Name, prod.Name, qty, prod.Price)

		prod.Stock -= qty
		c.Inventory = append(c.Inventory, &ClientInventory{
			Product:  *prod,
			Quantity: qty,
		})

		prod.Mutex.Unlock()
	}

	q <- c
}

func (c *Client) Leave(m *Market) {
	fmt.Printf("Client (%s) leaves\n", c.Name)

	m.Clients = slices.DeleteFunc(m.Clients, func(cl *Client) bool {
		return cl.ID == c.ID
	})
}
