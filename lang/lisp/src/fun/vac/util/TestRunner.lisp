(defpackage fun.vac.util.test-runner
  (:nicknames test-runner)
  (:use common-lisp)
  (:export
    parse-test
    ))

(in-package test-runner)

(defvar *test-runner-test-index* 0)

(defun parse-test (ok)
  (if ok
    (format t "✓ Test [~D] Passed~%" *test-runner-test-index*)
    (format t "✗ Test [~D] Failed~%" *test-runner-test-index*))

  (incf *test-runner-test-index*) nil)
