(import-module "fun/vac/util/RandomFactory")
(import-module "fun/vac/util/TestRunner")

(defun test-integer-n ()
  (random-factory:seed)

  (let (value)
    (dotimes (i 8192)
      (when (/= (random-factory:integer-n 0 0) 0)
        (return-from test-integer-n nil))

      (when (/= (random-factory:integer-n 1 1) 1)
        (return-from test-integer-n nil))

      (setf value (random-factory:integer-n 0 1))
      (when (or (< value 0) (< 1 value))
        (return-from test-integer-n nil))

      (setf value (random-factory:integer-n 100 10000))
      (when (/= (random-factory:integer-n value value) value)
        (return-from test-integer-n nil))
      (when (or (< value 100) (< 10000 value))
        (return-from test-integer-n nil))
      ))

  t)

(defun test-generate-even ()
  (random-factory:seed)

  (dotimes (i 8192)
    (unless (zerop (logand (random-factory:generate-even) 1))
      (return-from test-generate-even nil)))

  t)

(defun test-generate-odd ()
  (random-factory:seed)

  (dotimes (i 8192)
    (when (zerop (logand (random-factory:generate-odd) 1))
      (return-from test-generate-odd nil)))

  t)

(defun main ()

  (test-runner:parse-test (test-integer-n))

  (test-runner:parse-test (test-generate-even))

  (test-runner:parse-test (test-generate-odd)))
