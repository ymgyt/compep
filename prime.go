package compep

import (
	"math"
)

// IsPrime -
func IsPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

// GenPrimes -
func GenPrimes(n int) []int {
	ns := make([]int, 0, n)
	ns = append(ns, 2)

	m := 3
	for {
		if IsPrime(m) {
			ns = append(ns, m)
		}
		if len(ns) == n {
			break
		}
		m += 2
	}
	return ns[:n]
}
