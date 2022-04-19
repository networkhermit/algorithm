(defpackage muse.util.random-factory
  (:nicknames random-factory)
  (:use common-lisp)
  (:export
   seed
   gen-int-n
   gen-int
   gen-even
   gen-odd))

(in-package random-factory)

(defun seed ()
  (setf *random-state* (make-random-state t)) nil)

(defun gen-int-n (min max)
  (+ min (random (- max min -1))))

(defun gen-int ()
  (gen-int-n 1 2147483647))

(defun gen-even ()
  (ash (ash (gen-int) -1) 1))

(defun gen-odd ()
  (logior (gen-int) 1))
