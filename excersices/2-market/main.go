package main

import (
	"fmt"

	"github.com/ilyasa1211/go-exercises/excersices/2-market/entities"
	"github.com/ilyasa1211/go-exercises/excersices/2-market/factories"
	"github.com/ilyasa1211/go-exercises/excersices/2-market/services"
)

func main() {
	market := &entities.Market{}

	csFactory := &factories.CashierFactory{}
	prodFactory := &factories.ProductFactory{}
	clientFactory := &factories.ClientFactory{}

	market.AddCashiers(csFactory.Spawn(5))
	market.AddProducts(prodFactory.Spawn(10))
	market.AddClients(clientFactory.Spawn(100))

	mService := services.NewMarketService(market)

	mService.Run()

	fmt.Printf("Profits: %f\n", mService.GetProfits())
}
