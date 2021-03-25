package card
import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExampleTotal(){
	
	fmt.Println(Total([]types.Card{
		{
			Balance: 100000,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance:100000,
			Active:true,
		},
		{
				Balance: 200000,
				Active: true,
			},
	
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 100000,
			Active: false,
		},

	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -100000,
			Active: true,
		},
	}))
	// Output:
	//100000
	//300000
	//0
	//0
	
}
/*
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
*/