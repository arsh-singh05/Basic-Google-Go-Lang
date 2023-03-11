package main

import (
    "fmt"
)

func main() {
    // Input odds for each outcome
    var odds1, odds2 float64
    fmt.Print("Enter odds for outcome 1: ")
    fmt.Scan(&odds1)
    fmt.Print("Enter odds for outcome 2: ")
    fmt.Scan(&odds2)

    // Calculate implied probabilities
    p1 := 1.0 / odds1
    p2 := 1.0 / odds2

    // Calculate total payout and overround
    payout := 1.0 / (p1 + p2)
    overround := (p1 + p2 - 1.0) * 100.0

    // Check for arbitrage opportunity
    if payout > 1.0 {
        // Calculate stakes
        stake1 := (payout / p1) * 100.0
        stake2 := (payout / p2) * 100.0

        // Output results
        fmt.Println("Arbitrage opportunity found!")
        fmt.Printf("Stake on outcome 1: %f%%\n", stake1)
        fmt.Printf("Stake on outcome 2: %f%%\n", stake2)
        fmt.Printf("Total payout: %f%%\n", payout*100.0)
        fmt.Printf("Overround: %f%%\n", overround)
    } else {
        // Output results
        fmt.Println("No arbitrage opportunity found.")
        fmt.Printf("Total payout: %f%%\n", payout*100.0)
        fmt.Printf("Overround: %f%%\n", overround)
    }
}
