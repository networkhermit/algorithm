def iterative_procedure(n: int) -> int:
    sign = 1

    if n < 0:
        if (n & 1) == 0:
            sign = -1
        n = -n

    if n < 2:
        return n

    prev = 0
    curr = 1

    for _ in range(n - 1):
        prev, curr = curr, prev + curr

    return sign * curr


def recursive_procedure(n: int) -> int:
    if n < 0:
        val = recursive_procedure(-n)
        if (n & 1) == 0:
            return -val
        return val
    if n < 2:
        return n
    return recursive_procedure(n - 2) + recursive_procedure(n - 1)
