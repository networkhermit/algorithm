(defpackage fun.vac.util.random-factory
  (:nicknames random-factory)
  (:use common-lisp)
  (:export
    launch
    integer-n
    generate-integer
    generate-even
    generate-odd))

(in-package random-factory)

(defun launch ()
  (setf *random-state* (make-random-state t)) nil)

(defun integer-n (min max)
  (+ min (random (- max min))))

(defun generate-integer ()
  (integer-n 0 2147483647))

(defun generate-even ()
  (ash (ash (generate-integer) -1) 1))

(defun generate-odd ()
  (logior (generate-integer) 1))
