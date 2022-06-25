from muse.algorithms.number_theory import (
    coprimality,
    factorial,
    fibonacci_number,
    greatest_common_divisor,
    least_common_multiple,
    parity,
    primality,
    prime_sieves,
)


def is_coprime(m: int, n: int) -> bool:
    return coprimality.reduce_to_binary_gcd(m, n)


def fact(n: int) -> int:
    return factorial.iterative_procedure(n)


def fibonacci(n: int) -> int:
    return fibonacci_number.iterative_procedure(n)


def gcd(m: int, n: int) -> int:
    return greatest_common_divisor.iterative_binary_gcd(m, n)


def lcm(m: int, n: int) -> int:
    return least_common_multiple.reduce_to_binary_gcd(m, n)


def is_even(n: int) -> bool:
    return parity.bitwise_is_even(n)


def is_odd(n: int) -> bool:
    return parity.bitwise_is_odd(n)


def is_prime(n: int) -> bool:
    return primality.is_prime(n)


def is_composite(n: int) -> bool:
    return primality.is_composite(n)


def sieve_of_primes(n: int) -> list:
    return prime_sieves.sieve_of_eratosthenes(n)
