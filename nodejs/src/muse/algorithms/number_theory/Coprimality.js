"use strict"

const GreatestCommonDivisor = require("muse/algorithms/number_theory/GreatestCommonDivisor")

exports.reduceToBinaryGCD = (m, n) => {
    return GreatestCommonDivisor.iterativeBinaryGCD(m, n) === 1
}

exports.reduceToEuclidean = (m, n) => {
    return GreatestCommonDivisor.iterativeEuclidean(m, n) === 1
}
