package main

import "fmt"

// function that can take any number of int arguments
func sum(nums ...int) {
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println("Sum:", total)
}

func main() {
    sum(1, 2)
    sum(1, 2, 3)
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}

/*
This program uses a variadic function that can take any number of int arguments.
The ... before the type of the last parameter in the function signature indicates that the function can take any number of arguments of that type.
*/
