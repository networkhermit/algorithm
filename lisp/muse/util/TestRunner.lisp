(defpackage muse.util.test-runner
  (:nicknames test-runner)
  (:use common-lisp)
  (:export
    pick))

(in-package test-runner)

(defvar *test-runner-item-index* 0)

(defun pick (func)
  (if (funcall func)
    (format t "✓ Item [~D] PASSED~%" *test-runner-item-index*)
    (format t "✗ Item [~D] FAILED~%" *test-runner-item-index*))

  (incf *test-runner-item-index*))
