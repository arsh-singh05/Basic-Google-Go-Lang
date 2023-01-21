package main
import "fmt"

func apply(x int, f func(int) int) int {
    return f(x)
}

func square(x int) int {
    return x * x
}

func cube(x int) int {
    return x * x * x
}

func main() {
    resultsq := apply(5, square)
    resultcube := apply(5,cube)
    fmt.Println(resultsq)
    fmt.Println(resultcube)
}

/*
This program defines a function apply that takes an int argument x and a function f that takes an int argument and returns an int.
The apply function calls the function f with the argument x and returns the result.
The square function is defined as a function that takes an int argument and returns its square. The main function calls the apply function with the argument 5 and the square function, and assigns the result to the variable result and then it prints the value of result which is 25.
The cube function is defined as a function that takes an int argument and returns its cube. The main function calls the apply function with the argument 5 and the square function, and assigns the result to the variable result and then it prints the value of result which is 125.
/*