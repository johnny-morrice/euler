package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 10, "The maximum factor")
	flag.Parse()

	fmt.Println(Euler5(*n))
} 

func Euler5(top int) int {
	primes := primesunder(top)

	maxpows := make([]int, len(primes))

	for i := 2; i <= top; i++ {
		for j, p := range primes {
			rem := 0
			div := i
			pow := -1

			for ; rem == 0 ; pow++ {
				rem = div % p
				div = div / p
			}

			prev := &maxpows[j]
			
			if *prev < pow {
				*prev = pow
			}
		}
	}

	product := 1
	for i, p := range primes {
		pow := maxpows[i]

		product *= intpow(p, pow)
	}

	return product
}

func intpow(n int, m int) int {
	if m == 0 {
		return 1
	}

	pow := n
	for i := 0; i < m - 1; i++ {
		pow *= n
	}

	return pow
}

func primesunder(max int) []int {
	primes := []int{2, 3}

	for i := 4; i < max; i++ {
		isprime := true
		for _, p := range primes {
			if i % p == 0 {
				isprime = false
				break
			}
		}

		if isprime {
			primes = append(primes, i)
		}
	}

	return primes
}