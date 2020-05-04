(defpackage fun.vac.util.test-runner
  (:nicknames test-runner)
  (:use common-lisp)
  (:export
    parse-test))

(in-package test-runner)

(defvar *test-runner-item-index* 0)

(defun parse-test (ok)
  (if ok
    (format t "✓ Item [~D] PASSED~%" *test-runner-item-index*)
    (format t "✗ Item [~D] FAILED~%" *test-runner-item-index*))

  (incf *test-runner-item-index*))
