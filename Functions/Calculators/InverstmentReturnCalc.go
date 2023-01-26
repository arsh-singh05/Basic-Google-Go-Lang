package main

import (
    "fmt"
)

// InvestmentReturnCalculator takes in the initial investment amount, annual interest rate, and number of years and returns the final investment amount
func InvestmentReturnCalculator(initialInvestment float64, interestRate float64, numYears int) float64 {
    // formula for compound interest: A = P(1 + r/n)^nt
    finalInvestment := initialInvestment * (1 + interestRate/100)^(float64(numYears))
    return finalInvestment
}

func main() {
    initialInvestment := 1000.0
    interestRate := 5.0
    numYears := 10

    finalInvestment := InvestmentReturnCalculator(initialInvestment, interestRate, numYears)

    fmt.Printf("After %d years, an investment of $%.2f with an annual interest rate of %.2f%% will be worth $%.2f", 
    numYears, initialInvestment, interestRate, finalInvestment)
}

/*
This program defines a function called InvestmentReturnCalculator that takes in three inputs:
initialInvestment (float64): the amount invested
interestRate (float64): the annual interest rate, expressed as a decimal
numYears (int): the number of years the investment will be held
The function uses the formula for compound interest to calculate the final investment amount, which is returned by the function.
In the main function, the initial values for the investment, interest rate, and number of years are defined and passed as arguments to the InvestmentReturnCalculator function.
The final investment amount is then printed to the console using a formatted string.
*/
