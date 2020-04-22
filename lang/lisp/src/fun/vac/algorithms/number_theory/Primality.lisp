(defpackage fun.vac.algorithms.number-theory.primality
  (:nicknames primality)
  (:use common-lisp)
  (:export
    is-prime
    is-composite))

(in-package primality)

(defun is-prime (n)
  (when (< n 2)
    (return-from is-prime nil))
  (when (and (= (logand n 1) 0) (/= n 2))
    (return-from is-prime nil))

  (do ((bound (isqrt n))
       (i 3 (+ i 2)))
    ((> i bound) t)
    (when (= (rem n i) 0)
      (return-from is-prime nil))))

(defun is-composite (n)
  (and (> n 1) (not (is-prime n))))
