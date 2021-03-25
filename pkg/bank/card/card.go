package card
import "bank/pkg/bank/types"

func Total(cards []types.Card) types.Money {
	total := types.Money(0)

	for _, card := range cards {

		if !card.Active {
			continue
		}
		
		if card.Balance <=0 {
			continue
		}
		total +=card.Balance
		
		
	}
	return total
}

