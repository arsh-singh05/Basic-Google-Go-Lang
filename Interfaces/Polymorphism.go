package main

import "fmt"

type Shape interface {
    area() float64
}

type Rectangle struct {
    width  float64
    height float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func getArea(s Shape) float64 {
    return s.area()
}

func main() {
    r := Rectangle{width: 10, height: 20}
    c := Circle{radius: 5}
    fmt.Println("Area of Rectangle:", getArea(r))
    fmt.Println("Area of Circle:", getArea(c))
}

/*
This program defines an interface Shape that has a single method area() which is implemented by two structs Rectangle and Circle. 
The getArea function takes a value of type Shape as an argument and calls the area method on it. 
This allows getArea to work with any value that implements the Shape interface, regardless of its underlying type.
/*
