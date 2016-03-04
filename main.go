package main

import (
	"flag"
	"fmt"
	"goess/calc"
	"os"
	"strconv"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	areThereAnyArguments()
	estimations := stringSliceToFloatSlice(flag.Args())
	lambda := calc.Average(estimations)
	estimatedTime := calc.IntegratePoissonDistribution(lambda, 3*int(lambda)) * lambda
	fmt.Println("Estimated time: ", estimatedTime)
}

func stringSliceToFloatSlice(stringSlice []string) []float64 {
	var floatSlice []float64
	for _, arg := range stringSlice {
		number, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			if number < 0 {
				usage()
			}
			floatSlice = append(floatSlice, number)
		} else {
			usage()
		}
	}
	return floatSlice
}

func areThereAnyArguments() {
	if flag.NArg() == 0 {
		usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [estimated values...]\n example: goess 1 2 3\n\n only positive numbers are allowed!\n\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}
