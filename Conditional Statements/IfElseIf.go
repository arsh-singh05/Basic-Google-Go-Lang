package main
import "fmt"

func main() {
    x := 5
    if x > 0 {
        fmt.Println("x is positive")
    } else if x < 0 {
        fmt.Println("x is negative")
    } else {
        fmt.Println("x is zero")
    }
}