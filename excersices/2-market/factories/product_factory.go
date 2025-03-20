package factories

import (
	"sync"

	"github.com/brianvoe/gofakeit"
	"github.com/ilyasa1211/go-exercises/excersices/2-market/entities"
)

type ProductFactory struct{}

func (f *ProductFactory) Spawn(count uint) []*entities.Product {
	cs := make([]*entities.Product, count)

	for i := uint(0); i < count; i++ {
		cs[i] = &entities.Product{
			Name:  gofakeit.BeerName(),
			Price: gofakeit.Float64Range(2500, 20000),
			Stock: uint(gofakeit.Int32()),
			Mutex: &sync.Mutex{},
		}
	}

	return cs
}
