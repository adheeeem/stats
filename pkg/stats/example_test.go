package stats

import (
	"fmt"
	"github.com/adheeeem/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   50_00,
			Category: "korti milli",
			Status:   "FAIL",
		},
		{
			ID:       2,
			Amount:   10_000,
			Category: "visa",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   35_00,
			Category: "visa",
			Status:   "INPROGRESS",
		},
	}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 6750
}

func ExampleTotalInCategory() {
	category := types.Category("korti milli")
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   50_00,
			Category: "korti milli",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   10_000,
			Category: "visa",
			Status:   "FAIL",
		},
		{
			ID:       3,
			Amount:   35_00,
			Category: "visa",
			Status:   "OK",
		},
		{
			ID:       4,
			Amount:   15_000,
			Category: "korti milli",
			Status:   "INPROGRESS",
		},
	}
	result := TotalInCategory(payments, category)
	fmt.Println(result)
	// Output: 20000
}
