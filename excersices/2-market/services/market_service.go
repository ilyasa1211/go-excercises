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
		go c.PickupProducts(s.m, paymentQueue, &wg)
	}

	go func() {
		wg.Wait()
		close(paymentQueue)
	}()

	for i := 0; i < len(s.m.Cashiers); i++ {
		cashier := s.m.Cashiers[i]

		for c := range paymentQueue {
			wgCashiers.Add(1)
			go cashier.ServeClient(c, transactionLogs, s.m, &wgCashiers)
		}
	}

	go func() {
		wgCashiers.Wait()
		close(transactionLogs)
	}()

	for t := range transactionLogs {
		s.m.TransactionLogs = append(s.m.TransactionLogs, t)
	}

	if len(s.m.Clients) > 0 {
		fmt.Println("Retry failed transactions")
		s.Run()
	}
}

func (s *MarketService) GetProfits() float64 {
	var t float64

	fmt.Printf("Total transactions: %d\n", len(s.m.TransactionLogs))

	for _, log := range s.m.TransactionLogs {
		t += log.GetTotalPrice() - log.GetDiscount()
	}

	return t
}
