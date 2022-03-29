from muse.algorithms.number_theory import GreatestCommonDivisor


def reduceToBinaryGCD(m: int, n: int) -> bool:
    return GreatestCommonDivisor.iterativeBinaryGCD(m, n) == 1


def reduceToEuclidean(m: int, n: int) -> bool:
    return GreatestCommonDivisor.iterativeEuclidean(m, n) == 1
