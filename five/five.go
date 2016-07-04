package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 10, "The maximum factor")
	flag.Parse()

	fmt.Println(Euler5(int64(*n)))
} 

func Euler5(top int64) int64 {
	primes := primesunder(top)

	maxpows := make([]int64, len(primes))

	for i := int64(2); i <= top; i++ {
		for j, p := range primes {
			rem := int64(0)
			pow := int64(-1)
			div := i

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

	product := int64(1)
	for i, p := range primes {
		pow := maxpows[i]

		product *= intpow(p, pow)
	}

	return product
}

func intpow(n int64, m int64) int64 {
	if m == 0 {
		return 1
	}

	pow := n
	for i := int64(0); i < m - 1; i++ {
		pow *= n
	}

	return pow
}

func primesunder(max int64) []int64 {
	primes := []int64{2, 3}

	for i := int64(4); i < max; i++ {
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