package main

import (
    "fmt"
)

// Function to calculate GPA
func gpaCalculator(grades []float64) float64 {
    var total float64
    for _, grade := range grades {
        total += grade
    }
    return total / float64(len(grades))
}

func main() {
    grades := []float64{3.5, 4.0, 3.0, 2.5, 3.5}
    gpa := gpaCalculator(grades)
    fmt.Printf("GPA: %.2f", gpa)
}

/*
Here, the function gpaCalculator takes a slice of floats (grades) as an input, and returns a single float (GPA).
The function uses a for-range loop to iterate through the grades slice and add up the total grade points.
The total grade points are then divided by the number of grades to get the GPA, which is returned by the function.
The main function then calls the gpaCalculator function with a slice of grades, and stores the returned value in the variable gpa.
Finally, it prints the GPA with 2 decimal places using Printf.
*/
