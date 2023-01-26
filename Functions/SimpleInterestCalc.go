package main

import (
    "fmt"
)

// Function to calculate simple interest
func SimpleInterest(principal, rate, time float64) float64 {
    // Calculate simple interest using the formula:
    // interest = (principal * rate * time) / 100
    interest := (principal * rate * time) / 100
    return interest
}

func main() {
    // Declare variables for principal, rate and time
    var total,principal, rate, time float64

    // Get input for principal, rate and time
    fmt.Print("Enter principal: ")
    fmt.Scan(&principal)
    fmt.Print("Enter rate: ")
    fmt.Scan(&rate)
    fmt.Print("Enter time: ")
    fmt.Scan(&time)

    // Call the SimpleInterest function
    interest := SimpleInterest(principal, rate, time)
    total := principal+interest

    // Print the calculated interest
    fmt.Printf("Simple Interest: %.2f", interest)
    fmt.Printf("Total Amount: %.2f", total)
}

/*
This program calculates simple interest. 
It takes in 3 inputs from the user: the principal, the rate and the time.
The program defines a function SimpleInterest that takes in 3 arguments: the principal, the rate and the time. The formula for calculating simple interest is (principal * rate * time) / 100.
The function uses this formula to calculate the interest and then returns it.
In the main function, the program takes inputs from user and call the function by passing the inputs and then prints the result.
*/
