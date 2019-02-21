package FibonacciNumber

func IterativeProcedure(n int64) int64 {
    sign := int64(1)

    if n < 0 {
        if (n & 1) == 0 {
            sign = -1
        }
        n = -n
    }

    if n < 2 {
        return n
    }

    prev := int64(0)
    curr := int64(1)

    for i := int64(2); i <= n; i++ {
        prev, curr = curr, prev + curr
    }

    return sign * curr
}

func RecursiveProcedure(n int64) int64 {
    if n < 0 {
        if (n & 1) == 0 {
            return -RecursiveProcedure(-n)
        } else {
            return RecursiveProcedure(-n)
        }
    } else if n < 2 {
        return n
    } else {
        return RecursiveProcedure(n - 2) + RecursiveProcedure(n - 1)
    }
}
