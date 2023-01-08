package stats

import (
	"fmt"
	"github.com/adheeeem/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   50_00,
			Category: "korti milli",
		},
		{
			ID:       2,
			Amount:   10_000,
			Category: "visa",
		},
		{
			ID:       3,
			Amount:   35_00,
			Category: "visa",
		},
	}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 6166
}

func ExampleTotalInCategory() {
	category := types.Category("korti milli")
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   50_00,
			Category: "korti milli",
		},
		{
			ID:       2,
			Amount:   10_000,
			Category: "visa",
		},
		{
			ID:       3,
			Amount:   35_00,
			Category: "visa",
		},
		{
			ID:       4,
			Amount:   15_000,
			Category: "korti milli",
		},
	}
	result := TotalInCategory(payments, category)
	fmt.Println(result)
	// Output: 20000
}
