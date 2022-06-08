package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumNumbers[V Number](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers([]int64{1, 2, 4}),
		SumNumbers([]float64{1.1, 2.2, 4.4}))
}
