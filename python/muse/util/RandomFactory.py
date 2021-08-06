import random

def seed() -> None:
    random.seed()

def genIntN(min: int, max: int) -> int:
    return random.randint(min, max)

def genInt() -> int:
    return genIntN(0, 2_147_483_647)

def genEven() -> int:
    return genInt() >> 1 << 1

def genOdd() -> int:
    return genInt() | 1
