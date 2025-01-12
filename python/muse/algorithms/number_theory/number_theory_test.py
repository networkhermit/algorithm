from muse.algorithms import number_theory
from muse.algorithms.number_theory import (
    coprimality_test,
    factorial_test,
    fibonacci_number_test,
    greatest_common_divisor_test,
    least_common_multiple_test,
    parity_test,
    primality,
    primality_test,
    prime_sieves_test,
)
from muse.util import test_runner


def test_is_coprime() -> bool:
    return coprimality_test.derive(number_theory.is_coprime)()


def test_factorial() -> bool:
    return factorial_test.derive(number_theory.fact)()


def test_fibonacci() -> bool:
    return fibonacci_number_test.derive(number_theory.fibonacci)()


def test_gcd() -> bool:
    return greatest_common_divisor_test.derive(number_theory.gcd)()


def test_lcm() -> bool:
    return least_common_multiple_test.derive(number_theory.lcm)()


def test_is_even() -> bool:
    return parity_test.derive(number_theory.is_even, parity_test.EVEN)


def test_is_odd() -> bool:
    return parity_test.derive(number_theory.is_odd, parity_test.ODD)


def test_is_prime() -> bool:
    return primality_test.derive(number_theory.is_prime, primality_test.PRIME)


def test_is_composite() -> bool:
    return primality_test.derive(number_theory.is_composite, primality_test.COMPOSITE)


def test_sieve_of_primes() -> bool:
    return prime_sieves_test.derive(number_theory.sieve_of_primes, primality.is_prime)


def main() -> None:
    test_runner.pick(test_is_coprime)

    test_runner.pick(test_factorial)

    test_runner.pick(test_fibonacci)

    test_runner.pick(test_gcd)

    test_runner.pick(test_lcm)

    test_runner.pick(test_is_even)

    test_runner.pick(test_is_odd)

    test_runner.pick(test_is_prime)

    test_runner.pick(test_is_composite)

    test_runner.pick(test_sieve_of_primes)


if __name__ == "__main__":
    main()
