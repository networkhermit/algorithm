import random

def seed() -> None:
    random.seed()

def integerN(min: int, max: int) -> int:
    return random.randint(min, max)

def generateInteger() -> int:
    return integerN(0, 2_147_483_647)

def generateEven() -> int:
    return generateInteger() >> 1 << 1

def generateOdd() -> int:
    return generateInteger() | 1
