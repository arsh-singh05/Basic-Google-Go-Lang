package main

import (
    "fmt"
    "math"
)

// Function to calculate compound interest
func CompoundInterest(principal, rate float64, time int) float64 {
    // Calculate compound interest using the formula:
    // A = P(1 + r/n)^(nt)
    // where A = final amount, P = principal, r = rate, t = time and n = number of times compounded annually
    n := 12 // number of times compounded annually
    A := principal * math.Pow((1 + (rate / n)), (float64(time) * float64(n)))
    return A - principal
}

func main() {
    // Declare variables for principal, rate and time
    var principal, rate float64
    var time int

    // Get input for principal, rate and time
    fmt.Print("Enter principal: ")
    fmt.Scan(&principal)
    fmt.Print("Enter rate: ")
    fmt.Scan(&rate)
    fmt.Print("Enter time in years: ")
    fmt.Scan(&time)

    // Call the CompoundInterest function
    interest := CompoundInterest(principal, rate, time)

    // Print the calculated interest
    fmt.Printf("Compound Interest: %.2f", interest)
}

/*
This program calculates compound interest.
It takes in 3 inputs from the user: the principal, the rate and the time.
The program defines a function CompoundInterest that takes in 3 arguments: the principal, the rate and the time.
The formula for calculating compound interest is A = P(1 + r/n)^(nt) where A is the final amount, P is the principal, r is the rate, t is the time and n is the number of times compounded annually.
The function uses this formula to calculate the interest and then returns it.
In the main function, the program takes inputs from user and call the function by passing the inputs and then prints the result.
*/
