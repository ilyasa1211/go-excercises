package services

import (
	"fmt"
	"sync"

	"github.com/ilyasa1211/go-exercises/excersices/2-market/entities"
)

type MarketService struct {
	m *entities.Market
}

func NewMarketService(m *entities.Market) *MarketService {
	return &MarketService{m}
}

func (s *MarketService) Run() {
	paymentQueue := make(chan *entities.Client)
	transactionLogs := make(chan *entities.Transaction)

	var wg sync.WaitGroup
	var wgCashiers sync.WaitGroup

	for _, c := range s.m.Clients {
		wg.Add(1)
		go func(client *entities.Client) {
			defer wg.Done()
			c.PickupProducts(s.m, paymentQueue)
		}(c)
	}

	go func() {
		wg.Wait()
		close(paymentQueue)
	}()

	for i := 0; i < len(s.m.Cashiers); i++ {
		wgCashiers.Add(1)

		go func(cashier *entities.Cashier) {
			defer wgCashiers.Done()

			for c := range paymentQueue {
				cashier.ServeClient(c, transactionLogs)
			}
		}(s.m.Cashiers[i])
	}

	go func() {
		wgCashiers.Wait()
		close(transactionLogs)
	}()

	for t := range transactionLogs {
		s.m.TransactionLogs = append(s.m.TransactionLogs, t)
	}
}

func (s *MarketService) GetProfits() float64 {
	var t float64

	fmt.Printf("Total transactions: %d\n", len(s.m.TransactionLogs))

	for _, log := range s.m.TransactionLogs {
		t += log.GetTotalPrice()
	}

	return t
}
