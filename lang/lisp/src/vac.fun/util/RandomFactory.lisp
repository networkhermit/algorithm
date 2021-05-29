(defpackage fun.vac.util.random-factory
  (:nicknames random-factory)
  (:use common-lisp)
  (:export
    seed
    integer-n
    generate-integer
    generate-even
    generate-odd))

(in-package random-factory)

(defun seed ()
  (setf *random-state* (make-random-state t)) nil)

(defun integer-n (min max)
  (+ min (random (- max min -1))))

(defun generate-integer ()
  (integer-n 0 2147483647))

(defun generate-even ()
  (ash (ash (generate-integer) -1) 1))

(defun generate-odd ()
  (logior (generate-integer) 1))
