package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // 120
}

/*
This program defines a function factorial that takes an int argument n and uses recursion to calculate the factorial of n.
The function uses a base case of n == 0 and returns 1 in that case. 
In other cases, it calls itself with the argument n-1 and multiplies the result by n
The main function calls the factorial function with the argument 5 and prints the result which is 120.
/*