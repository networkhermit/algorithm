(defpackage muse.util.test-runner
  (:nicknames test-runner)
  (:use common-lisp)
  (:export
   pick))

(in-package test-runner)

(defvar *item-index* 0)

(defun pick (func)
  (if (funcall func)
      (format t "✓ Item [~D] PASSED~%" *item-index*)
      (format t "✗ Item [~D] FAILED~%" *item-index*))

  (incf *item-index*))
