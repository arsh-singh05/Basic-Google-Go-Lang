package main
import "fmt"

func main() 
{
    x := 5
    y := 3
    z := 4

    if x > 0 {
        if y > 0 {
            if z > 0 {
                fmt.Println("x, y, and z are all positive")
            } else {
                fmt.Println("x and y are positive but z is non-positive")
            }
        } else {
            fmt.Println("x is positive but y is non-positive")
        }
    } else {
        fmt.Println("x is non-positive")
    }
}

/*
This program declares three variables x, y, and z with the values 5, 3, and 4 respectively, and uses nested if-else statements to check if each of the variables is greater than 0. 
If all the variables are greater than 0, the program prints "x, y, and z are all positive".
If any of the variables is less than or equal to 0, the program will print the corresponding message.
/*