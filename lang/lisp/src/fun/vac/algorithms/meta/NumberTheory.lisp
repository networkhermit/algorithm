(import-module "fun/vac/algorithms/number_theory/Coprimality")
(import-module "fun/vac/algorithms/number_theory/Factorial")
(import-module "fun/vac/algorithms/number_theory/FibonacciNumber")
(import-module "fun/vac/algorithms/number_theory/GreatestCommonDivisor")
(import-module "fun/vac/algorithms/number_theory/LeastCommonMultiple")
(import-module "fun/vac/algorithms/number_theory/Parity")
(import-module "fun/vac/algorithms/number_theory/Primality")

(defpackage fun.vac.algorithms.meta.number-theory
  (:nicknames number-theory)
  (:use common-lisp)
  (:shadow gcd lcm)
  (:export
    is-coprime
    factorial
    fibonacci
    gcd
    lcm
    is-even
    is-odd
    is-prime
    is-composite
    ))

(in-package number-theory)

(defun is-coprime (m n)
  (coprimality::reduce-to-binary-gcd m n))

(defun factorial (n)
  (factorial::iterative-procedure n))

(defun fibonacci (n)
  (fibonacci-number::iterative-procedure n))

(defun gcd (m n)
  (greatest-common-divisor::iterative-binary-gcd m n))

(defun lcm (m n)
  (least-common-multiple::reduce-to-binary-gcd m n))

(defun is-even (n)
  (parity::bitwise-is-even n))

(defun is-odd (n)
  (parity::bitwise-is-odd n))

(defun is-prime (n)
  (primality::is-prime n))

(defun is-composite (n)
  (primality::is-composite n))
