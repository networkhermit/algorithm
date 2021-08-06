(import-module "muse/util/TestRunner")

(defun test-pick ()
  (dotimes (i 10)
    (test-runner:pick (lambda () (/= (logand i 1) 0)))))

(defun main ()
  (test-pick))
