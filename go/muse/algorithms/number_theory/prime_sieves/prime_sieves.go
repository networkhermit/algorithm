package prime_sieves

import "math"

func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	size := (n + 1) >> 1

	arr := make([]bool, size)

	numPrimes := size
	for i, bound := 3, int(math.Sqrt(float64(n)))+1; i < bound; i += 2 {
		if arr[i>>1] {
			continue
		}
		for j := i * i; j <= n; j += i << 1 {
			if !arr[j>>1] {
				arr[j>>1] = true
				numPrimes--
			}
		}
	}

	primes := make([]int, 0, numPrimes)

	primes = append(primes, 2)

	for i := 3; i <= n; i += 2 {
		if !arr[i>>1] {
			primes = append(primes, i)
		}
	}

	return primes
}
