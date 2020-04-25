(import-module "fun/vac/util/RandomFactory")
(import-module "fun/vac/util/TestRunner")

(defun test-generate-even ()
  (random-factory:launch)

  (let (value)
    (dotimes (i 8192)
      (setf value (random-factory:generate-even))
      (when (/= (logand value 1) 0)
        (return nil))))

  t)

(defun test-generate-odd ()
  (random-factory:launch)

  (let (value)
    (dotimes (i 8192)
      (setf value (random-factory:generate-odd))
      (when (zerop (logand value 1))
        (return nil))))

  t)

(defun main ()

  (test-runner:parse-test (test-generate-even))

  (test-runner:parse-test (test-generate-odd)))
