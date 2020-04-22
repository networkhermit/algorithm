(defpackage fun.vac.algorithms.number-theory.factorial
  (:nicknames factorial)
  (:use common-lisp)
  (:export
    iterative-procedure
    recursive-procedure))

(in-package factorial)

(defun iterative-procedure (n)
  (when (< n 0)
    (return-from iterative-procedure 0))

  (do ((result 1)
       (i 1 (1+ i)))
    ((> i n) result)
    (setf result (* result i))))

(defun recursive-procedure (n)
  (when (< n 0)
    (return-from recursive-procedure 0))

  (if (= n 0)
    1
    (* (recursive-procedure (- n 1)) n)))
