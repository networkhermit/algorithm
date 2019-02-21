"use strict"

exports.launch = () => {
    // preserve for consistent interface
}

exports.integerN = (min, max) => {
    return min + Math.floor(Math.random() * (max - min))
}

exports.generateInteger = () => {
    return this.integerN(0, 2147483647)
}

exports.generateEven = () => {
    return this.generateInteger() >>> 1 << 1
}

exports.generateOdd = () => {
    return this.generateInteger() | 1
}
