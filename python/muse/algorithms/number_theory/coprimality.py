from muse.algorithms.number_theory import greatest_common_divisor


def reduce_to_binary_gcd(m: int, n: int) -> bool:
    return greatest_common_divisor.iterative_binary_gcd(m, n) == 1


def reduce_to_euclidean(m: int, n: int) -> bool:
    return greatest_common_divisor.iterative_euclidean(m, n) == 1
