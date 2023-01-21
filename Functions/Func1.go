package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

func main() {
    result := sum(1, 2, 3, 4, 5)
    fmt.Println(result)
}

/*
This program defines a function sum that takes a variable number of int arguments, using the ... syntax.
The function uses a loop to sum up all the numbers and returns the total as an int.
The main function calls the sum function with the arguments 1, 2, 3, 4, 5, and assigns the result to the variable result, then it prints the value of result which is 15.
/*