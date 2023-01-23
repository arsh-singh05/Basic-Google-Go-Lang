package main

import "fmt"

func sierpinski(n int) {
    for i := 0; i < n; i++ {
        for j := 0; j < (1 << uint(n-i))-1; j++ {
            fmt.Print(" ")
        }
        for j := 0; j < (1 << uint(i))-1; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}

func main() {
    sierpinski(5)
}

/*
The sierpinski function takes an integer as an argument representing the depth of the triangle.
The outer loop iterates from 0 to n to generate the n rows of the triangle.
The first inner loop generates the spaces before the stars in each row. The number of spaces is calculated by (1 << uint(n-i))-1 where n is the depth of the triangle and i is the current row.
The second inner loop generates the stars in each row. The number of stars is calculated by (1 << uint(i))-1 where i is the current row.
In main function, we call the sierpinski function and pass the depth of the triangle as an argument.
This code will produce an ASCII representation of a Sierpinski triangle with a depth of 5. 
To change the depth of the triangle, you can simply change the argument passed to the sierpinski function in the main function.
*/
