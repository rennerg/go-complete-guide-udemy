package main

import (
	"fmt"
	"math"
)

// Investment represents an investment with its parameters.
type Investment struct {
	Principal        float64 // Initial amount of money
	AnnualRate       float64 // Annual interest rate (in percentage)
	Years            int     // Number of years the money is invested
	CompoundsPerYear int     // Number of times interest is compounded per year
}

// FutureValue calculates the future value of the investment.
func (inv Investment) FutureValue() float64 {
	ratePerPeriod := inv.AnnualRate / 100 / float64(inv.CompoundsPerYear)
	totalPeriods := inv.Years * inv.CompoundsPerYear
	futureValue := inv.Principal * math.Pow(1+ratePerPeriod, float64(totalPeriods))
	return futureValue
}

func main() {
	investment := getUserInput()
	futureValue := investment.FutureValue()
	fmt.Printf("Future Value of the Investment: $%.2f\n", futureValue)
}

func getUserInput() Investment {
	var principal float64
	var annualRate float64
	var years int
	var compoundsPerYear int

	fmt.Print("Enter the principal amount: ")
	fmt.Scanln(&principal)

	fmt.Print("Enter the annual interest rate (in %): ")
	fmt.Scanln(&annualRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scanln(&years)

	fmt.Print("Enter the number of times interest is compounded per year: ")
	fmt.Scanln(&compoundsPerYear)

	return Investment{
		Principal:        principal,
		AnnualRate:       annualRate,
		Years:            years,
		CompoundsPerYear: compoundsPerYear,
	}
}
