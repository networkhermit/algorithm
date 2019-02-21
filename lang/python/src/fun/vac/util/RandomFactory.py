import random

def launch() -> None:
    random.seed()

def integerN(min: int, max: int) -> int:
    return random.randint(min, max - 1)

def generateInteger() -> int:
    return integerN(0, 2_147_483_647)

def generateEven() -> int:
    return generateInteger() >> 1 << 1

def generateOdd() -> int:
    return generateInteger() | 1
