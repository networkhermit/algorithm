(use "muse/algorithms/number-theory/factorial")
(use "muse/util/test-runner")

(defun test-factorial ()
  (let ((sample (make-array '(13 2) :initial-contents
                            '(( 0         1)
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
                              (12 479001600)))))

    (dotimes (i (array-dimension sample 0))
      (when (/= (factorial:iterative-procedure (aref sample i 0)) (aref sample i 1))
        (return-from test-factorial nil)))

    (dotimes (i (array-dimension sample 0))
      (when (/= (factorial:recursive-procedure (aref sample i 0)) (aref sample i 1))
        (return-from test-factorial nil))))

  t)

(defun main ()
  (test-runner:pick #'test-factorial))
