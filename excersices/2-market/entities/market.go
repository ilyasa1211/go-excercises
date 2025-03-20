package entities

type Market struct {
	Inventories     []*Product
	TransactionLogs []*Transaction
	Cashiers        []*Cashier
	Clients         []*Client
}

func (m *Market) AddProducts(p []*Product) {
	m.Inventories = append(m.Inventories, p...)
}
func (m *Market) AddCashiers(c []*Cashier) {
	m.Cashiers = append(m.Cashiers, c...)
}
func (m *Market) AddClients(c []*Client) {
	m.Clients = append(m.Clients, c...)
}
func (m *Market) AddTransactions(t []*Transaction) {
	m.TransactionLogs = append(m.TransactionLogs, t...)
}
