package factories

import (
	"github.com/brianvoe/gofakeit"
	"github.com/ilyasa1211/go-exercises/excersices/2-market/entities"
)

type ClientFactory struct{}

func (f *ClientFactory) Spawn(count uint) []*entities.Client {
	cs := make([]*entities.Client, count)

	for i := uint(0); i < count; i++ {
		cs[i] = &entities.Client{
			User: entities.User{
				Name: gofakeit.Name(),
			},
		}
	}

	return cs
}
