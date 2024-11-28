package main

import (
	"fmt"
	"math"
)

func main() {
    investmentAmount, years, expectedReturnRate, inflationRate := 1000.0, 10.0, 5.5, 2.5
    
    fmt.Print("Enter inflation rate: ")
    fmt.Scan(&inflationRate)

    fv, frv := logic(investmentAmount, expectedReturnRate, inflationRate, years)

    formattedFV := fmt.Sprintf("Future value: %.1f\n", fv)
    formattedRFV := fmt.Sprintf("Adjusted for inflation: %.1f\n", frv)

    fmt.Print(formattedFV, formattedRFV)
    // fmt.Printf("Future value: %.1f\nAdjusted for inflation: %.1f\n", futureValue, futureRealValue)
    // back ticks usage
    fmt.Printf(`
    Future value: %.1f
    Adjusted for inflation: %.1f
    `, fv, frv)
}

func logic(investmentAmount, expectedReturnRate, inflationRate, years float64) (fv, frv float64) {
    fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
    frv = fv / math.Pow(1 + inflationRate / 100, years)
    return
}