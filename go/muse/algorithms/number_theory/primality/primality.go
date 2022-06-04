package primality

import "math"

func IsPrime(n int64) bool {
	if n < 2 {
		return false
	}
	if (n&1) == 0 && n != 2 {
		return false
	}

	for i, bound := int64(3), int64(math.Sqrt(float64(n)))+1; i < bound; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsComposite(n int64) bool {
	return n > 1 && !IsPrime(n)
}
