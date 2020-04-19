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

  (let ((result 1))
    (do ((i 1 (1+ i)))
      ((> i n))
      (setf result (* result i)))

    (return-from iterative-procedure result)))

(defun recursive-procedure (n)
  (when (< n 0)
    (return-from recursive-procedure 0))

  (if (= n 0)
    1
    (* (recursive-procedure (- n 1)) n)))
