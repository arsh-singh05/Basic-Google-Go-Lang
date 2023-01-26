package main

import (
  "fmt"
)

func mortgageCalculator(principal float64, interestRate float64, termInYears int) float64 {
    interestRate = interestRate / 100 / 12
    termInMonths := termInYears * 12
    monthlyPayment := principal * interestRate / (1 - (1 / (1 + interestRate)^termInMonths))
    return monthlyPayment
}

func main() {
    principal := 250000.00
    interestRate := 3.5
    termInYears := 30
    fmt.Printf("Your monthly mortgage payment is: $%.2f", mortgageCalculator(principal, interestRate, termInYears))
}

/*
This function takes in three parameters: the principal amount of the mortgage, the annual interest rate, and the term of the mortgage in years.
It first calculates the monthly interest rate and the term in months, and then uses these values to calculate the monthly mortgage payment using the standard mortgage formula.
The function then returns this value, which is printed to the screen in the main function.
*/
