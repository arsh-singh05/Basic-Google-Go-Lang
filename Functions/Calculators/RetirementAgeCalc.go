package main

import (
	"fmt"
)

func retirementAgeCalculator(currentAge, retirementAge int) {
	fmt.Println("Your current age is:", currentAge)
	fmt.Println("Your retirement age is:", retirementAge)
	fmt.Println("You have", retirementAge-currentAge, "years left for retirement.")
}

func main() {
	var currentAge int
	var retirementAge int
	fmt.Print("Enter your current age: ")
	fmt.Scan(&currentAge)
	fmt.Print("Enter your retirement age: ")
	fmt.Scan(&retirementAge)
	retirementAgeCalculator(currentAge, retirementAge)
}

/*
This program calculates the number of years left for the user's retirement.
It takes two inputs: currentAge and retirementAge.
The retirementAgeCalculator function takes in these two inputs and calculates the number of years left for retirement.
The main function takes input from the user for currentAge and retirementAge and calls the retirementAgeCalculator function.
It then prints out the user's current age, retirement age, and the number of years left for retirement.
*/
