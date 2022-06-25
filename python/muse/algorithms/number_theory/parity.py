def modulo_is_even(n: int) -> bool:
    return n % 2 == 0


def modulo_is_odd(n: int) -> bool:
    return n % 2 != 0


def bitwise_is_even(n: int) -> bool:
    return (n & 1) == 0


def bitwise_is_odd(n: int) -> bool:
    return (n & 1) != 0
