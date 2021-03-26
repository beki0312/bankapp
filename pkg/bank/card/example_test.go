package card
import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExamplePaymentSource(){
	var NewCard types.PaymentSource
	payments := [] types.Card{
	{
		PAN: "10",
		Balance:500,
		Active:true,
	},
	{
			PAN: "11",
			Balance: 100,
			Active: true,
		},
	{
		PAN: "12",
		Balance: 500,
		Active: false,
	},
	}
	maxes :=PaymentSources(payments)
	for _, max := range maxes {
		NewCard.Number=max.Number
	fmt.Println(NewCard.Number)
	}
	//Output:
	//10
	//11
}
