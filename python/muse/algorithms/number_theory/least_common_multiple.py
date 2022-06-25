from muse.algorithms.number_theory import greatest_common_divisor


def reduce_to_binary_gcd(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    if m == 0 or n == 0:
        return 0
    return m // greatest_common_divisor.iterative_binary_gcd(m, n) * n


def reduce_to_euclidean(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    if m == 0 or n == 0:
        return 0
    return m // greatest_common_divisor.iterative_euclidean(m, n) * n
