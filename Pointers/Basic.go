package main
import "fmt"

func main()
{
    x := 5
    var p *int = &x
    *p = 10
    fmt.Println(x)
}

// This program declares a variable x with the value 5, creates a pointer p to x, then it assigns the value 10 to the memory location the pointer is pointing at
