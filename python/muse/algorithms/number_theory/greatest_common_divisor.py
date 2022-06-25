def iterative_binary_gcd(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    shift = 0

    while True:
        if m == n:
            return m << shift
        if m == 0:
            return n << shift
        if n == 0:
            return m << shift

        if (m & 1) == 0:
            m >>= 1
            if (n & 1) == 0:
                n >>= 1
                shift += 1
        elif (n & 1) == 0:
            n >>= 1
        elif m > n:
            m = (m - n) >> 1
        else:
            n = (n - m) >> 1


def recursive_binary_gcd(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    if m == n:
        return m
    if m == 0:
        return n
    if n == 0:
        return m

    if (m & 1) == 0:
        if (n & 1) == 0:
            return recursive_binary_gcd(m >> 1, n >> 1) << 1
        return recursive_binary_gcd(m >> 1, n)
    if (n & 1) == 0:
        return recursive_binary_gcd(m, n >> 1)
    if m > n:
        return recursive_binary_gcd((m - n) >> 1, n)
    return recursive_binary_gcd(m, (n - m) >> 1)


def iterative_euclidean(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    while n != 0:
        m, n = n, m % n

    return m


def recursive_euclidean(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    if n == 0:
        return m
    return recursive_euclidean(n, m % n)
