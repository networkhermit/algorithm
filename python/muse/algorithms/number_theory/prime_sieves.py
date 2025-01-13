import math


def sieve_of_eratosthenes(n: int) -> list[int]:
    if n < 2:
        return []

    size = (n + 1) >> 1

    arr = [False] * size

    num_primes = size
    for i in range(3, math.isqrt(n) + 1, 2):
        if arr[i >> 1]:
            continue
        for j in range(i * i, n + 1, i << 1):
            if not arr[j >> 1]:
                arr[j >> 1] = True
                num_primes -= 1

    primes = [0] * num_primes

    primes[0] = 2

    curr = 1
    for i in range(3, n + 1, 2):
        if not arr[i >> 1]:
            primes[curr] = i
            curr += 1

    return primes
