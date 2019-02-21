(defpackage fun.vac.algorithms.number-theory.primality
  (:nicknames primality)
  (:use common-lisp)
  (:export
    is-prime
    is-composite
    ))

(in-package primality)

(defun is-prime (n)
  (when (< n 2)
    (return-from is-prime nil))
  (when (and (= (logand n 1) 0) (/= n 2))
    (return-from is-prime nil))

  (let ((bound (isqrt n)))
    (do ((i 3 (+ i 2)))
      ((> i bound))
      (when (= (rem n i) 0)
        (return-from is-prime nil))))

  t)

(defun is-composite (n)
  (and (> n 1) (not (is-prime n))))
