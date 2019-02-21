(import-module "fun/vac/algorithms/number_theory/Factorial")
(import-module "fun/vac/util/TestRunner")

(defun test-factorial ()
  (let ((mapping (make-array '(13 2) :initial-contents
                             '(
                               ( 0         1)
                               ( 1         1)
                               ( 2         2)
                               ( 3         6)
                               ( 4        24)
                               ( 5       120)
                               ( 6       720)
                               ( 7      5040)
                               ( 8     40320)
                               ( 9    362880)
                               (10   3628800)
                               (11  39916800)
                               (12 479001600)
                               ))))

    (let ((instances (array-dimension mapping 0)))

      (dotimes (i instances)
        (when (/= (factorial:iterative-procedure (aref mapping i 0)) (aref mapping i 1))
          (return-from test-factorial nil)))

      (dotimes (i instances)
        (when (/= (factorial:recursive-procedure (aref mapping i 0)) (aref mapping i 1))
          (return-from test-factorial nil)))))
  t)

(defun main ()
  (test-runner:parse-test (test-factorial)))
