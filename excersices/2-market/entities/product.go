package entities

import "sync"

type Product struct {
	Name  string
	Price float64
	Stock uint
	Mutex *sync.Mutex
}
