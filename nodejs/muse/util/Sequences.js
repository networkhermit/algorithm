"use strict"

const QuickSort     = require("muse/algorithms/sorting/QuickSort")
const RandomFactory = require("muse/util/RandomFactory")

exports.inspect = (arr) => {
    console.log("[")
    for (let i = 0, length = arr.length; i < length; i++) {
        process.stdout.write(`\t#${i.toString(16).toUpperCase().padStart(4, "0")}  -->  ${arr[i]}\n`)
    }
    console.log("]")
}

exports.isSorted = (arr) => {
    for (let i = 1, length = arr.length; i < length; i++) {
        if (arr[i] < arr[i - 1]) {
            return false
        }
    }

    return true
}

exports.parityChecksum = (arr) => {
    let checksum = 0

    for (let v of arr) {
        checksum ^= v
    }

    return checksum
}

exports.reverse = (arr) => {
    let k
    let temp

    for (let i = 0, bound = Math.floor(arr.length >>> 1), length = arr.length; i < bound; i++) {
        k = length - i - 1
        temp = arr[i]
        arr[i] = arr[k]
        arr[k] = temp
    }
}

exports.shuffle = (arr) => {
    let k
    let temp

    RandomFactory.seed()
    for (let i = 0, length = arr.length; i < length; i++) {
        k = RandomFactory.genIntN(i, length)
        temp = arr[i]
        arr[i] = arr[k]
        arr[k] = temp
    }
}

exports.sort = (arr) => {
    QuickSort.sort(arr)
}
