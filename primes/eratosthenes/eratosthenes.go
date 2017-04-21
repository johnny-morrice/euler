package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 100, "Max")
	flag.Parse()

	primes := eratosthenes(*n)

	for _, p := range primes {
		fmt.Println(p)
	}
}

func eratosthenes(maxPrime int) []int {
	count := maxPrime - 2
	maxPossible := maxPrime / 2
	mark := make([]bool, count)

	for p := 2; p < maxPossible; {
		for m := 2; (p * m) - 2 < count; m++ {
			q := p * m
			mark[q - 2] = true
		}

		for i := p + 1; i < count; i++ {
			if !mark[i - 2] {
				p = i
				break
			}
		}
	}

	primes := []int{}
	for i := 0; i < count; i++ {
		if !mark[i] {
			primes = append(primes, i + 2)
		}
	}

	return primes
}
