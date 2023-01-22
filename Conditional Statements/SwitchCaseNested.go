package main
import "fmt"

func main() {
    x := 5
    y := 3

OuterLoop:
    switch x {
    case 0:
        fmt.Println("x is zero")
    case 1, 2:
        fmt.Println("x is one or two")
        switch y {
        case 0:
            fmt.Println("y is zero")
            break OuterLoop
        case 1, 2:
            fmt.Println("y is one or two")
        default:
            fmt.Println("y is something else")
        }
    default:
        fmt.Println("x is something else")
    }
}

/*
This program declares two variables x and y with the values 5 and 3 respectively, and uses a nested switch statement with a label statement.
The label statement allows breaking out of multiple levels of nested loops or switch statements by specifying the label of the outer loop or switch statement.
In this example, when the value of y is 0, it will break out of the outer switch statement and the program will only print "y is zero".
/*