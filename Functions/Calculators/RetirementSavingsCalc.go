package main

import (
    "fmt"
)

// RetirementSavingsCalculator function calculates the savings needed for retirement
// based on the current age, retirement age, desired annual income, and annual return on investment.
func RetirementSavingsCalculator(currentAge int, retirementAge int, desiredIncome float64, annualROI float64) float64 {
    // calculate the number of years until retirement
    yearsUntilRetirement := retirementAge - currentAge

    // calculate the desired income at retirement, accounting for inflation
    desiredIncomeAtRetirement := desiredIncome * (1 + 0.05)^yearsUntilRetirement

    // calculate the savings needed at retirement using the following formula:
    // savings = desired income at retirement / (annual ROI / 100)
    savings := desiredIncomeAtRetirement / (annualROI / 100)

    return savings
}

func main() {
    currentAge := 30
    retirementAge := 65
    desiredIncome := 100000.0
    annualROI := 6.0

    savings := RetirementSavingsCalculator(currentAge, retirementAge, desiredIncome, annualROI)

    fmt.Printf("You need to save $%.2f for your retirement", savings)
}

/*
In this example, the function RetirementSavingsCalculator takes in 4 parameters, the current age, retirement age, desired annual income and annual return on investment.
It then calculates the number of years until retirement and uses that to calculate the desired income at retirement by accounting for 5% inflation (global average).
Then it calculates the savings needed at retirement by dividing desired income at retirement by annualROI/100.
The main function calls this RetirementSavingsCalculator function and prints the savings needed for retirement.
It's important to note that this is a simplified example and in real world scenario more factors would be considered while calculating savings needed for retirement.
*/
