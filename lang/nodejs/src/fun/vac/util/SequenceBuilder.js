"use strict"

const RandomFactory = require("fun/vac/util/RandomFactory")
const Sequences     = require("fun/vac/util/Sequences")

exports.packIncreasing = (arr) => {
    RandomFactory.launch()
    arr[0] = RandomFactory.integerN(1, 4)
    for (let i = 1, length = arr.length; i < length; i++) {
        arr[i] = arr[i - 1] + RandomFactory.integerN(1, 4)
    }
}

exports.packRandom = (arr) => {
    RandomFactory.launch()
    for (let i = 0, length = arr.length; i < length; i++) {
        arr[i] = RandomFactory.generateInteger()
    }
}

exports.packDecreasing = (arr) => {
    this.packIncreasing(arr)
    Sequences.reverse(arr)
}
