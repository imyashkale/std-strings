package mapgo

import (
	"fmt"
	"github.com/imyashkale/std-strings/models"
	"time"
)

func CalculatePorpolioValuation(portpolio ...int) int {
	total := 0
	for _, investment := range portpolio {
		total += investment
	}
	return total
}

func MapGo() {

	profileInfo := models.ProfileInfo{
		FirstName:  "Yash",
		LastName:   "Kale",
		DOB:        "07-07-99",
		City:       "Aurangabad",
		FavActress: "Sabrina",
		Age:        21,
		CreatedAt:  time.Now().Month(),
		NetWorth: models.NetWorth{
			Stocks: []string{
				"HDFCBANK",
				"ITC",
				"LT",
			},
			Bonds: []string{
				"SBIBOND",
				"ICICIBOND",
			},
			Cash: 1200000,
		},
	}

	stock := make(map[int8]string)

	for index, item := range profileInfo.NetWorth.Stocks {
		stock[int8(index)] = item
		fmt.Printf("[%v] : %v \n", index, item)
	}

	for k, v := range stock {
		fmt.Printf("Key : %v Value %v \n", k, v)
	}

	st, ok := stock[7]
	if ok {
		fmt.Printf("%v Found In Portfolio \n", st)
	} else {
		fmt.Printf("%v Not Found In Portfolio \n", st)
	}

	// Accepting Inifinite number of Argument
	valuation := CalculatePorpolioValuation([]int{100, 200, 300, 400, 4000}...)
	fmt.Printf("Calculated Value : %v \n", valuation)

	first, last := profileInfo.UpperCaseName()

	fmt.Printf("In Caps : %v %v \n", first, last)

	models.CalculateTaxOnCash(&profileInfo)

	fmt.Printf("After Taxed : %d \n", profileInfo.Tax)
}
