package main

import "fmt"

// Function to calculate credit card interest
func creditCardInterest(balance, interestRate, minimumPayment float64) float64 {
    //Calculating the interest
    interest := balance * (interestRate / 100)
    //Calculating the new balance with added interest
    newBalance := balance + interest
    //Checking if the minimum payment is greater than the interest
    if minimumPayment > interest {
        //Returning the new balance minus the minimum payment
        return newBalance - minimumPayment
    }
    //Returning the new balance
    return newBalance
}

func main() {
    //Initial balance
    balance := 1000.0
    //Interest rate
    interestRate := 20.0
    //Minimum payment
    minimumPayment := 50.0
    //Calling the creditCardInterest function
    remainingBalance := creditCardInterest(balance, interestRate, minimumPayment)
    //Printing the remaining balance
    fmt.Println("Remaining Balance: ", remainingBalance)
}

/*
In the above program, we have defined a function creditCardInterest() which takes three parameters balance, interestRate and minimumPayment which are all of float64 data type.
The function calculates the interest on the given balance using the given interest rate, then it adds that interest to the balance to get the new balance.
It then checks if the minimum payment is greater than the interest, if it is, it subtracts the minimum payment from the new balance to get the remaining balance.
If the minimum payment is not greater than the interest, the function returns the new balance.
In the main function, we have defined the initial balance, interest rate, and minimum payment and calling the creditCardInterest function and printing the remaining balance.
*/
