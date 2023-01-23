package main

import "fmt"

var memo = make(map[int]int)

func fib(n int) int {
    if n <= 0 {
        return 0
    } else if n == 1 {
        return 1
    } else if _, ok := memo[n]; ok {
        return memo[n]
    }
    memo[n] = fib(n-1) + fib(n-2)
    return memo[n]
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(fib(i))
    }
}

/*
This program demonstrates how to use memoization to optimize a recursive function.
In this example, the fibonacci function is implemented recursively, but it uses a map to store already calculated results and returns it directly instead of recalculating it.
*/
