package main
import "fmt"

func main() {
    x := 5
    switch {
    case x < 0:
        fmt.Println("x is negative")
    case x > 0 && x < 10:
        fmt.Println("x is between 0 and 10")
    case x >= 10:
        fmt.Println("x is greater or equal to 10")
    }
}