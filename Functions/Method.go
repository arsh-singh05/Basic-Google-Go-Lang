package main

import "fmt"

type Point struct {
    x, y int
}

// Method
func (p *Point) Scale(factor int) {
    p.x *= factor
    p.y *= factor
}

func main() {
    point := Point{3, 4}
    point.Scale(2)
    fmt.Println(point)
}

/*
This program demonstrates how to create a method in Go. A method is a function that is bound to a specific type, in this case, the Point struct. 
The Scale method takes a factor and multiplies the x and y fields of the Point struct by that factor.
In Go, the first parameter of a method is always the receiver, which is a pointer to the struct the method is bound to, in this case, the *Point pointer.
*/
