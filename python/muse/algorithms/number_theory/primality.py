import math


def is_prime(n: int) -> bool:
    if n < 2:
        return False
    if (n & 1) == 0 and n != 2:
        return False

    for i in range(3, math.isqrt(n) + 1, 2):
        if n % i == 0:
            return False

    return True


def is_composite(n: int) -> bool:
    return n > 1 and not is_prime(n)
