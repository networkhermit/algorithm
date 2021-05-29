(import-module "vac.fun/algorithms/number_theory/GreatestCommonDivisor")

(defpackage fun.vac.algorithms.number-theory.least-common-multiple
  (:nicknames least-common-multiple)
  (:use common-lisp)
  (:export
    reduce-to-binary-gcd
    reduce-to-euclidean))

(in-package least-common-multiple)

(defun reduce-to-binary-gcd (m n)
  (when (minusp m)
    (setf m (- m)))
  (when (minusp n)
    (setf n (- n)))

  (if (or (zerop m) (zerop n))
    0
    (* (truncate m (greatest-common-divisor::iterative-binary-gcd m n)) n)))

(defun reduce-to-euclidean (m n)
  (when (minusp m)
    (setf m (- m)))
  (when (minusp n)
    (setf n (- n)))

  (if (or (zerop m) (zerop n))
    0
    (* (truncate m (greatest-common-divisor::iterative-euclidean m n)) n)))
