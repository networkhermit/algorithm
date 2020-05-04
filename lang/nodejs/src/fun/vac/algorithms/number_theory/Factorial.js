"use strict"

exports.iterativeProcedure = (n) => {
    if (n < 0) {
        return 0
    }

    let result = 1
    for (let i = 1; i <= n; i++) {
        result *= i
    }

    return result
}

exports.recursiveProcedure = (n) => {
    if (n < 0) {
        return 0
    }

    if (n === 0) {
        return 1
    }
    return this.recursiveProcedure(n - 1) * n
}
