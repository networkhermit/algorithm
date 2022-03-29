package GreatestCommonDivisor

func IterativeBinaryGCD(m int64, n int64) int64 {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	shift := uint(0)

	for {
		if m == n {
			return m << shift
		}
		if m == 0 {
			return n << shift
		}
		if n == 0 {
			return m << shift
		}

		if (m & 1) == 0 {
			m >>= 1
			if (n & 1) == 0 {
				n >>= 1
				shift++
			}
		} else if (n & 1) == 0 {
			n >>= 1
		} else if m > n {
			m = (m - n) >> 1
		} else {
			n = (n - m) >> 1
		}
	}
}

func RecursiveBinaryGCD(m int64, n int64) int64 {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	if m == n {
		return m
	}
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	if (m & 1) == 0 {
		if (n & 1) == 0 {
			return RecursiveBinaryGCD(m>>1, n>>1) << 1
		}
		return RecursiveBinaryGCD(m>>1, n)
	}
	if (n & 1) == 0 {
		return RecursiveBinaryGCD(m, n>>1)
	}
	if m > n {
		return RecursiveBinaryGCD((m-n)>>1, n)
	}
	return RecursiveBinaryGCD(m, (n-m)>>1)
}

func IterativeEuclidean(m int64, n int64) int64 {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	for n != 0 {
		m, n = n, m%n
	}

	return m
}

func RecursiveEuclidean(m int64, n int64) int64 {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return m
	}
	return RecursiveEuclidean(n, m%n)
}
