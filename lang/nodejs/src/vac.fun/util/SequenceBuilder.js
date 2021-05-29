"use strict"

const RandomFactory = require("vac.fun/util/RandomFactory")
const Sequences     = require("vac.fun/util/Sequences")

exports.packIncreasing = (arr) => {
    RandomFactory.seed()
    arr[0] = RandomFactory.integerN(1, 3)
    for (let i = 1, length = arr.length; i < length; i++) {
        arr[i] = arr[i - 1] + RandomFactory.integerN(1, 3)
    }
}

exports.packRandom = (arr) => {
    RandomFactory.seed()
    for (let i = 0, length = arr.length; i < length; i++) {
        arr[i] = RandomFactory.generateInteger()
    }
}

exports.packDecreasing = (arr) => {
    this.packIncreasing(arr)
    Sequences.reverse(arr)
}
