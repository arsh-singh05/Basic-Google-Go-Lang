package main
import "fmt"

func main() 
{
    var x int = 2
    switch x 
    {
    case 1:
        fmt.Println("x is 1")
    case 2:
        fmt.Println("x is 2")
    default:
        fmt.Println("x is not 1 or 2")
    }
}

// This program uses a switch-case statement to check the value of x. If x is 1, it prints "x is 1" to the console. If x is 2, it prints "x is 2" to the console. Otherwise, it prints "x is not 1 or 2".