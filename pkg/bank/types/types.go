package types
//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирами и т.д.)
type Money int64

//Currency представляет код валюты
type Currency string

//коды валют
const(
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)
//PAN представляет номер карты
type PAN string

//Card представляет  информацию о платёжном карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	MinBalance Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}


type Payment struct{
	ID int
	Amount Money
}
