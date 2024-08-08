package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 5.08
	var investmentAmount, years, expectedReturnRate float64

	fmt.Println("Enter investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Println("Enter return rate")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Enter investment tenure")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
