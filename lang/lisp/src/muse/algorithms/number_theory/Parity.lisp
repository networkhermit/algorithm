(defpackage muse.algorithms.number-theory.parity
  (:nicknames parity)
  (:use common-lisp)
  (:export
    modulo-is-even
    modulo-is-odd
    bitwise-is-even
    bitwise-is-odd))

(in-package parity)

(defun modulo-is-even (n)
  (zerop (rem n 2)))

(defun modulo-is-odd (n)
  (/= (rem n 2) 0))

(defun bitwise-is-even (n)
  (zerop (logand n 1)))

(defun bitwise-is-odd (n)
  (/= (logand n 1) 0))
