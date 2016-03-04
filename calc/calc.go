package calc

import (
	"math"
	"math/big"
)

func IntegratePoissonDistribution(lambda float64, k int) float64 {
	if k == 0 {
		return 0.0
	}
	return PoissonDistribution(lambda, k) + IntegratePoissonDistribution(lambda, k-1)
}

func PoissonDistribution(lambda float64, k int) float64 {
	counter := (math.Pow(lambda, float64(k))) * (math.Exp((lambda * -1)))
	denominator := Factorial(k)
	bigCounter := big.NewFloat(counter)
	//fmt.Println("Counter: ", bigCounter)
	//fmt.Println("Denominator: ", denominator)
	poisson := bigCounter.Quo(bigCounter, denominator)
	floatPoisson, _ := poisson.Float64()
	return floatPoisson
}

func Average(numbers []float64) float64 {
	average := 0.0
	for _, number := range numbers {
		average += number
	}
	return average / float64(len(numbers))
}

func Factorial(number int) *big.Float {
	if number == 0 {
		return big.NewFloat(1.0)
	}
	bigNumber := big.NewFloat(float64(number))
	return bigNumber.Mul(bigNumber, Factorial(number-1))
}
