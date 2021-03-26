package card
import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card)[]types.PaymentSource{
	var operations []types.PaymentSource
	for _, card := range cards {
		if card.Balance > 0 && card.Active == true {
		operations=append(operations, types.PaymentSource{ Type:"card", Number:card.PAN, Balance:card.Balance})
		}
	}
	return operations
}

