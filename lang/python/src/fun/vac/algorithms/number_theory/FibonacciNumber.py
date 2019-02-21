def iterativeProcedure(n: int) -> int:
    sign = 1

    if n < 0:
        if (n & 1) == 0:
            sign = -1
        n = -n

    if n < 2:
        return n

    prev = 0
    curr = 1

    for i in range(2, n + 1):
        prev, curr = curr, prev + curr

    return sign * curr

def recursiveProcedure(n: int) -> int:
    if n < 0:
        if (n & 1) == 0:
            return -recursiveProcedure(-n)
        else:
            return recursiveProcedure(-n)
    elif n < 2:
        return n
    else:
        return recursiveProcedure(n - 2) + recursiveProcedure(n - 1)
