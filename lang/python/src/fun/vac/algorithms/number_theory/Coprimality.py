from fun.vac.algorithms.number_theory import GreatestCommonDivisor

def reduceToBinaryGCD(m: int, n: int) -> int:
    return GreatestCommonDivisor.iterativeBinaryGCD(m, n) == 1

def reduceToEuclidean(m: int, n: int) -> int:
    return GreatestCommonDivisor.iterativeEuclidean(m, n) == 1
