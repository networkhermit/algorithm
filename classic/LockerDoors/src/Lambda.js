"use strict"

const bruteForce = (n) => {
    const arr = new Array(n)

    arr.fill(false)

    for (let i = 1; i <= n; i++) {
        for (let j = i - 1; j < n; j += i) {
            arr[j] = !arr[j]
        }
    }

    return arr
}

const process = (n) => {
    const arr = Array(n)

    let temp = 0
    for (let i = 1; i <= n; i++) {
        temp = Math.floor(Math.sqrt(i))
        arr[i - 1] = temp * temp === i
    }

    return arr
}

if (module === require.main) {
    const N = 100

    const arr = process(N)

    for (let i = 0; i < N; i++) {
        if (arr[i]) {
            console.log("%d\t  OPEN", i + 1)
        }
    }
}
