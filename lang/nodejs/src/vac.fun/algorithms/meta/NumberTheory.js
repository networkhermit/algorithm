"use strict"

const Coprimality           = require("vac.fun/algorithms/number_theory/Coprimality")
const Factorial             = require("vac.fun/algorithms/number_theory/Factorial")
const FibonacciNumber       = require("vac.fun/algorithms/number_theory/FibonacciNumber")
const GreatestCommonDivisor = require("vac.fun/algorithms/number_theory/GreatestCommonDivisor")
const LeastCommonMultiple   = require("vac.fun/algorithms/number_theory/LeastCommonMultiple")
const Parity                = require("vac.fun/algorithms/number_theory/Parity")
const Primality             = require("vac.fun/algorithms/number_theory/Primality")
const PrimeSieves           = require("vac.fun/algorithms/number_theory/PrimeSieves")

exports.isCoprime = (m, n) => {
    return Coprimality.reduceToBinaryGCD(m, n)
}

exports.factorial = (n) => {
    return Factorial.iterativeProcedure(n)
}

exports.fibonacci = (n) => {
    return FibonacciNumber.iterativeProcedure(n)
}

exports.gcd = (m, n) => {
    return GreatestCommonDivisor.iterativeBinaryGCD(m, n)
}

exports.lcm = (m, n) => {
    return LeastCommonMultiple.reduceToBinaryGCD(m, n)
}

exports.isEven = (n) => {
    return Parity.bitwiseIsEven(n)
}

exports.isOdd = (n) => {
    return Parity.bitwiseIsOdd(n)
}

exports.isPrime = (n) => {
    return Primality.isPrime(n)
}

exports.isComposite = (n) => {
    return Primality.isComposite(n)
}

exports.sieveOfPrimes = (n) => {
    return PrimeSieves.sieveOfEratosthenes(n)
}
