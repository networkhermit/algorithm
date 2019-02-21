def iterativeBinaryGCD(m: int, n: int) -> int:
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
            if (n & 1) == 0:
                m >>= 1
                n >>= 1
                shift += 1
            else:
                m >>= 1
        else:
            if (n & 1) == 0:
                n >>= 1
            else:
                if m > n:
                    m = (m - n) >> 1
                else:
                    n = (n - m) >> 1

def recursiveBinaryGCD(m: int, n: int) -> int:
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
            return recursiveBinaryGCD(m >> 1, n >> 1) << 1
        else:
            return recursiveBinaryGCD(m >> 1, n)
    else:
        if (n & 1) == 0:
            return recursiveBinaryGCD(m, n >> 1)
        else:
            if m > n:
                return recursiveBinaryGCD((m - n) >> 1, n)
            else:
                return recursiveBinaryGCD(m, (n - m) >> 1)

def iterativeEuclidean(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    while n != 0:
        m, n = n, m % n

    return m

def recursiveEuclidean(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    if n == 0:
        return m
    else:
        return recursiveEuclidean(n, m % n)
