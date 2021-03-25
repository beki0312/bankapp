package card
import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExampleTotal(){
	cards := [] types.Card{
	{
		Balance:500,
		Active:true,
	},
	{
			Balance: -1,
			Active: true,
		},
	{
		Balance: 500,
		Active: false,
	},

	}


	fmt.Println(Total(cards))
//output: 500
}
