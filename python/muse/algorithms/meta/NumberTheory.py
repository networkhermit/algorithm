from muse.algorithms.number_theory import Coprimality
from muse.algorithms.number_theory import Factorial
from muse.algorithms.number_theory import FibonacciNumber
from muse.algorithms.number_theory import GreatestCommonDivisor
from muse.algorithms.number_theory import LeastCommonMultiple
from muse.algorithms.number_theory import Parity
from muse.algorithms.number_theory import Primality
from muse.algorithms.number_theory import PrimeSieves


def isCoprime(m: int, n: int) -> bool:
    return Coprimality.reduceToBinaryGCD(m, n)


def factorial(n: int) -> int:
    return Factorial.iterativeProcedure(n)


def fibonacci(n: int) -> int:
    return FibonacciNumber.iterativeProcedure(n)


def gcd(m: int, n: int) -> int:
    return GreatestCommonDivisor.iterativeBinaryGCD(m, n)


def lcm(m: int, n: int) -> int:
    return LeastCommonMultiple.reduceToBinaryGCD(m, n)


def isEven(n: int) -> bool:
    return Parity.bitwiseIsEven(n)


def isOdd(n: int) -> bool:
    return Parity.bitwiseIsOdd(n)


def isPrime(n: int) -> bool:
    return Primality.isPrime(n)


def isComposite(n: int) -> bool:
    return Primality.isComposite(n)


def sieveOfPrimes(n: int) -> list:
    return PrimeSieves.sieveOfEratosthenes(n)
