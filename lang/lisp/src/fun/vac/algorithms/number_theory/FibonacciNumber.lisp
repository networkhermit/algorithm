(defpackage fun.vac.algorithms.number-theory.fibonacci-number
  (:nicknames fibonacci-number)
  (:use common-lisp)
  (:export
    iterative-procedure
    recursive-procedure))

(in-package fibonacci-number)

(defun iterative-procedure (n)
  (let ((sign 1))

    (when (< n 0)
      (when (= (logand n 1) 0)
        (setf sign -1))
      (setf n (- n)))

    (when (< n 2)
      (return-from iterative-procedure n))

    (do ((prev 0)
         (curr 1)

         next
         (i 2 (1+ i)))
      ((> i n) (* sign curr))
      (setf next (+ prev curr))
      (shiftf prev curr next))))

(defun recursive-procedure (n)
  (cond ((< n 0)
         (if (= (logand n 1) 0)
           (- (recursive-procedure (- n)))
           (recursive-procedure (- n))))
        ((< n 2)
         n)
        (t
          (+ (recursive-procedure (- n 2)) (recursive-procedure (- n 1))))))
