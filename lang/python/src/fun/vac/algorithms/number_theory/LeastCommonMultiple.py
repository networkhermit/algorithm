from fun.vac.algorithms.number_theory import GreatestCommonDivisor

def reduceToBinaryGCD(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    if m == 0 or n == 0:
        return 0
    else:
        return m // GreatestCommonDivisor.iterativeBinaryGCD(m, n) * n

def reduceToEuclidean(m: int, n: int) -> int:
    if m < 0:
        m = -m
    if n < 0:
        n = -n

    if m == 0 or n == 0:
        return 0
    else:
        return m // GreatestCommonDivisor.iterativeEuclidean(m, n) * n
