"use strict"

exports.iterativeProcedure = (n) => {
    let sign = 1

    if (n < 0) {
        if ((n & 1) === 0) {
            sign = -1
        }
        n = -n
    }

    if (n < 2) {
        return n
    }

    let prev = 0
    let curr = 1

    let next = 0
    for (let i = 2; i <= n; i++) {
        next = prev + curr
        prev = curr
        curr = next
    }

    return sign * curr
}

exports.recursiveProcedure = (n) => {
    if (n < 0) {
        if ((n & 1) === 0) {
            return -this.recursiveProcedure(-n)
        } else {
            return this.recursiveProcedure(-n)
        }
    } else if (n < 2) {
        return n
    } else {
        return this.recursiveProcedure(n - 2) + this.recursiveProcedure(n - 1)
    }
}
