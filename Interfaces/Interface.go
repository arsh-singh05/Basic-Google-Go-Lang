package main
import "fmt"

type Shape interface 
{
    area() float64
}

type Rectangle struct 
{
    width, height float64
}

func (r Rectangle) area() float64 
{
    return r.width * r.height
}

func getArea(s Shape) float64 
{
    return s.area()
}

func main()
{
    r := Rectangle{width: 10, height: 5}
    fmt.Println("Area of rectangle:", getArea(r))
}

/* 
This program uses an interface Shape that has a single method area() and a struct Rectangle that implements the Shape interface.
The getArea function takes a Shape interface as an argument and calls the area() method on it.
The main function creates a Rectangle struct and passes it to the getArea function, which then calculates and prints the area of the rectangle.
*/
