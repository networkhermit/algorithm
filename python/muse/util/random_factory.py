import secrets


def gen_int_n(min: int, max: int) -> int:
    return min + secrets.randbelow(max - min + 1)


def gen_int() -> int:
    return gen_int_n(1, 2_147_483_647)


def gen_even() -> int:
    return gen_int() >> 1 << 1


def gen_odd() -> int:
    return gen_int() | 1
