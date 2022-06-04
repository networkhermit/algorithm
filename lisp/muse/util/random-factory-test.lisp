(use "muse/util/random-factory")
(use "muse/util/test-runner")

(defun test-gen-int-n ()
  (random-factory:seed)

  (let (value)
    (dotimes (i 8192)
      (when (/= (random-factory:gen-int-n 0 0) 0)
        (return-from test-gen-int-n nil))

      (when (/= (random-factory:gen-int-n 1 1) 1)
        (return-from test-gen-int-n nil))

      (setf value (random-factory:gen-int-n 0 1))
      (when (or (< value 0) (> value 1))
        (return-from test-gen-int-n nil))

      (setf value (random-factory:gen-int-n 100 10000))
      (when (/= (random-factory:gen-int-n value value) value)
        (return-from test-gen-int-n nil))
      (when (or (< value 100) (> value 10000))
        (return-from test-gen-int-n nil))
      ))

  t)

(defun test-gen-even ()
  (random-factory:seed)

  (dotimes (i 8192)
    (unless (zerop (logand (random-factory:gen-even) 1))
      (return-from test-gen-even nil)))

  t)

(defun test-gen-odd ()
  (random-factory:seed)

  (dotimes (i 8192)
    (when (zerop (logand (random-factory:gen-odd) 1))
      (return-from test-gen-odd nil)))

  t)

(defun main ()
  (test-runner:pick #'test-gen-int-n)

  (test-runner:pick #'test-gen-even)

  (test-runner:pick #'test-gen-odd))
