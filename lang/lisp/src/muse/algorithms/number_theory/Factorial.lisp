(defpackage muse.algorithms.number-theory.factorial
  (:nicknames factorial)
  (:use common-lisp)
  (:export
    iterative-procedure
    recursive-procedure))

(in-package factorial)

(defun iterative-procedure (n)
  (when (minusp n)
    (return-from iterative-procedure 0))

  (do ((result 1)
       (i 1 (1+ i)))
    ((> i n) result)
    (setf result (* result i))))

(defun recursive-procedure (n)
  (when (minusp n)
    (return-from recursive-procedure 0))

  (if (zerop n)
    1
    (* (recursive-procedure (- n 1)) n)))
