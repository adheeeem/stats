package stats

import "github.com/adheeeem/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	length := 0
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
			length++
		}
	}
	sum /= types.Money(length)
	return sum
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	total := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			total += payment.Amount
		}
	}
	return total
}
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	number := map[types.Category]int{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		number[payment.Category]++
	}
	for category := range categories {
		categories[category] /= types.Money(number[category])
	}
	return categories
}
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for category, money := range second {
		result[category] = money - first[category]
	}
	return result
}
