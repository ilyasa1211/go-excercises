package factories

import (
	"sync"

	"github.com/brianvoe/gofakeit"
	"github.com/ilyasa1211/go-exercises/excersices/2-market/entities"
)

type CashierFactory struct{}

func (f *CashierFactory) Spawn(count uint) []*entities.Cashier {
	cs := make([]*entities.Cashier, count)

	for i := uint(0); i < count; i++ {
		cs[i] = &entities.Cashier{
			User: entities.User{
				Name: gofakeit.Name(),
			},
			Mutex: &sync.Mutex{},
		}
	}

	return cs
}
