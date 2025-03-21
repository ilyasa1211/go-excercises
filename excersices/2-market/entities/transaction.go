package entities

import "time"

type Transaction struct {
	Items []*ClientInventory
	Date  time.Time
}

func (t *Transaction) GetTotalPrice() float64 {
	var res float64

	for _, p := range t.Items {
		res += p.Price * float64(p.Quantity)
	}

	return res
}

func (t *Transaction) GetDiscount() float64 {
	itemCountDiscount := 3
	discountPercentage := 10. / 100.

	if len(t.Items) > itemCountDiscount {
		return float64(t.GetTotalPrice() * discountPercentage)
	}

	return float64(0)
}
