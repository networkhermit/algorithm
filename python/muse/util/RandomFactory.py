import secrets

def genIntN(min: int, max: int) -> int:
    return min + secrets.randbelow(max - min + 1)

def genInt() -> int:
    return genIntN(1, 2_147_483_647)

def genEven() -> int:
    return genInt() >> 1 << 1

def genOdd() -> int:
    return genInt() | 1
