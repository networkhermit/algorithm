(import-module "muse/util/TestRunner")

(defun test-pick ()
  (dotimes (i 9)
    (test-runner:pick (lambda () (/= (logand i 1) 0)))))

(defun main ()
  (test-runner:pick #'test-pick))
