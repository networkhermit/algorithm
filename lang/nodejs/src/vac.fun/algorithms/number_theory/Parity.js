"use strict"

exports.moduloIsEven = (n) => {
    return n % 2 === 0
}

exports.moduloIsOdd = (n) => {
    return n % 2 !== 0
}

exports.bitwiseIsEven = (n) => {
    return (n & 1) === 0
}

exports.bitwiseIsOdd = (n) => {
    return (n & 1) !== 0
}
