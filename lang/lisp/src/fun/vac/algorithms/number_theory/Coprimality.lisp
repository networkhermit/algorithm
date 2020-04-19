(import-module "fun/vac/algorithms/number_theory/GreatestCommonDivisor")

(defpackage fun.vac.algorithms.number-theory.coprimality
  (:nicknames coprimality)
  (:use common-lisp)
  (:export
    reduce-to-binary-gcd
    reduce-to-euclidean))

(in-package coprimality)

(defun reduce-to-binary-gcd (m n)
  (= (greatest-common-divisor:iterative-binary-gcd m n) 1))

(defun reduce-to-euclidean (m n)
  (= (greatest-common-divisor:iterative-euclidean m n) 1))
