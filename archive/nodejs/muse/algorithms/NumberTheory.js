import * as Coprimality from './number_theory/Coprimality.js';
import * as Factorial from './number_theory/Factorial.js';
import * as FibonacciNumber from './number_theory/FibonacciNumber.js';
import * as GreatestCommonDivisor from './number_theory/GreatestCommonDivisor.js';
import * as LeastCommonMultiple from './number_theory/LeastCommonMultiple.js';
import * as Parity from './number_theory/Parity.js';
import * as Primality from './number_theory/Primality.js';
import * as PrimeSieves from './number_theory/PrimeSieves.js';

export const isCoprime = (m, n) => {
  return Coprimality.reduceToBinaryGCD(m, n);
};

export const factorial = (n) => {
  return Factorial.iterativeProcedure(n);
};

export const fibonacci = (n) => {
  return FibonacciNumber.iterativeProcedure(n);
};

export const gcd = (m, n) => {
  return GreatestCommonDivisor.iterativeBinaryGCD(m, n);
};

export const lcm = (m, n) => {
  return LeastCommonMultiple.reduceToBinaryGCD(m, n);
};

export const isEven = (n) => {
  return Parity.bitwiseIsEven(n);
};

export const isOdd = (n) => {
  return Parity.bitwiseIsOdd(n);
};

export const isPrime = (n) => {
  return Primality.isPrime(n);
};

export const isComposite = (n) => {
  return Primality.isComposite(n);
};

export const sieveOfPrimes = (n) => {
  return PrimeSieves.sieveOfEratosthenes(n);
};
