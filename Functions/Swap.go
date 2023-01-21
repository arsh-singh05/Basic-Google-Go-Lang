package main
import "fmt"

func swap(x, y int) (int, int) {
    return y, x
}

func main() {
    a, b := 3, 4
    a, b = swap(a, b)
    fmt.Println(a, b)
}

/*
This program defines a function swap that takes two int arguments x and y and returns them in reverse order.
The main function declares two variables a and b with the values 3 and 4, respectively, then it calls the swap function with the variables a and b, and assigns the returned values to the same variables, swapping their values.
The program then prints the values of a and b which are 4 and 3 respectively.
/*