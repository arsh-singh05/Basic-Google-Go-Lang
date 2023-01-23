package main

import "fmt"

// Currency exchange rates
var rates = map[string]map[string]float64{
	"USD": {"USD": 1.0, "EUR": 0.8, "GBP": 0.6},
	"EUR": {"USD": 1.25, "EUR": 1.0, "GBP": 0.75},
	"GBP": {"USD": 1.67, "EUR": 1.33, "GBP": 1.0},
}

// Finds the maximum profit by buying and selling currencies
func arbitrage(amount float64, currencies []string) float64 {
	profit := 0.0
	for i := 0; i < len(currencies)-1; i++ {
		profit += amount * (rates[currencies[i]][currencies[i+1]] - 1)
		amount *= rates[currencies[i]][currencies[i+1]]
	}
	return profit
}

func main() {
	profit := arbitrage(100, []string{"USD", "EUR", "GBP", "USD"})
	fmt.Printf("Maximum profit: $%.2f", profit)
}

/*
The arbitrage function takes two arguments: an initial amount of money, and a slice of strings representing the currencies to buy and sell.
The function uses a map of exchange rates to calculate the profit made by buying and selling currencies. The key of the map represents the currency being sold, and the value is another map, representing the currency being bought.
The function uses a for loop to iterate over the slice of currencies, starting with the first currency and ending with the last currency.
Inside the loop, the function calculates the profit by multiplying the amount by the exchange rate, and subtracting 1.
The function keeps track of the total profit and the amount of money in each iteration.
In main function, we call the arbitrage function and pass the amount of money, and the currencies to buy and sell.
This code will calculate the maximum profit that can be made by buying and selling currencies according to the exchange rates provided, and will print the profit.
*/
