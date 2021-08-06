"use strict"

exports.seed = () => {
    // preserve for consistent interface
}

exports.genIntN = (min, max) => {
    return min + Math.floor(Math.random() * (max - min))
}

exports.genInt = () => {
    return this.genIntN(0, 2147483647)
}

exports.genEven = () => {
    return this.genInt() >>> 1 << 1
}

exports.genOdd = () => {
    return this.genInt() | 1
}
