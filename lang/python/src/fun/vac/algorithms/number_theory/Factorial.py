def iterativeProcedure(n: int) -> int:
    if n < 0:
        return 0

    result = 1
    for i in range(1, n + 1):
        result *= i

    return result


def recursiveProcedure(n: int) -> int:
    if n < 0:
        return 0

    if n == 0:
        return 1
    else:
        return recursiveProcedure(n - 1) * n
