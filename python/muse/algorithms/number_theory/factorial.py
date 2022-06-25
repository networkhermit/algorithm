def iterative_procedure(n: int) -> int:
    if n < 0:
        return 0

    result = 1
    for i in range(1, n + 1):
        result *= i

    return result


def recursive_procedure(n: int) -> int:
    if n < 0:
        return 0

    if n == 0:
        return 1
    return recursive_procedure(n - 1) * n
